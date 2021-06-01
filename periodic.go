package health

import (
	"time"
)

type periodic struct {
	fn   func() error
	d    time.Duration
	done chan struct{}
}

func Periodic(fn func() error, d time.Duration) Check {
	return &periodic{fn, d, make(chan struct{})}
}

func (p *periodic) Notify(ch chan<- error) {
	go func() {
	L:
		for {
			t := time.NewTimer(p.d)
			select {
			case <-t.C:
			case <-p.done:
				if !t.Stop() {
					<-t.C
				}
				break L
			}

			ch <- p.fn()
		}
	}()
}

func (p *periodic) Stop(ch chan<- error) {
}
