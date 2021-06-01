package health

import (
	"io"
	"net/http"
)

type Monitor struct {
	services []service
	done     chan struct{}
}

func New(opts ...Option) *Monitor {
	m := Monitor{}
	for _, o := range opts {
		o.Apply(&m)
	}
	return &m
}

func (m *Monitor) Start() {
	m.done = make(chan struct{})
	for _, s := range m.services {
		go s.loop(m.done)
	}
}

func (m *Monitor) Stop() {
	close(m.done)
}

func (m *Monitor) Status() Status {
	for _, s := range m.services {
		s.m.RLock()
		if s.status != Healthy {
			return s.status
		}
		s.m.RUnlock()
	}
	return Healthy
}

func (m *Monitor) Endpoint() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, m.Status().GoString())
	})
}
