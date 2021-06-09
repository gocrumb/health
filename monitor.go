package health

import (
	"io"
	"log"
	"net/http"
	"sync"
)

type Monitor struct {
	checks     []Check
	states     map[Check]State
	status     State
	trs        chan transaction
	wg         sync.WaitGroup
	stop, done chan struct{}
	failfns    []func()
}

func New(opts ...Option) *Monitor {
	m := Monitor{
		states: map[Check]State{},
		trs:    make(chan transaction),
	}
	for _, o := range opts {
		o.Apply(&m)
	}
	return &m
}

func (m *Monitor) Start() {
	log.Print("Starting")

	m.stop = make(chan struct{})
	m.done = make(chan struct{})
	go m.loop()

	for _, c := range m.checks {
		m.wg.Add(1)
		go func(c Check) {
			defer m.wg.Done()

			ch := make(chan error)
			c.Notify(ch)
			defer c.Stop(ch)

			var state State = Starting
			m.trs <- stateChange{c, Stopped, Starting}

		L:
			for {
				select {
				case err := <-ch:
					switch state {
					case Starting:
						if err == nil {
							state = Healthy
							m.trs <- stateChange{c, Starting, Healthy}
						}

					case Healthy:
						if err != nil {
							state = Unhealthy
							m.trs <- stateChange{c, Healthy, Unhealthy}
						}

					case Unhealthy:
						if err == nil {
							state = Healthy
							m.trs <- stateChange{c, Unhealthy, Healthy}
						}
					}

				case <-m.stop:
					state = Stopped
					m.trs <- stateChange{c, state, Stopped}
					break L
				}
			}
		}(c)
	}
}

func (m *Monitor) Stop() {
	log.Print("Stopping")

	close(m.stop)
	m.wg.Wait()
	close(m.done)
}

func (m *Monitor) loop() {
	for {
		select {
		case tr := <-m.trs:
			switch tr := tr.(type) {
			case stateChange:
				m.states[tr.c] = tr.new

				status := Healthy
				for _, s := range m.states {
					if s != Healthy {
						status = s
						break
					}
				}
				if m.status != status {
					log.Printf("%s -> %s", m.status, status)
					m.status = status
				}

				if tr.new == Unhealthy {
					go func() {
						for _, fn := range m.failfns {
							fn()
						}
					}()
				}
			}

		case <-m.done:
		}
	}
}

func (m *Monitor) State(c Check) State {
	return m.states[c]
}

func (m *Monitor) Status() State {
	return m.status
}

func (m *Monitor) Endpoint() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, m.Status().String())
	})
}
