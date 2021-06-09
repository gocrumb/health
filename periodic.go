package health

import (
	"sync"
	"time"
)

type periodic struct {
	fn   func() error
	d    time.Duration
	chs  map[chan<- error]bool
	done chan struct{}
	m    sync.Mutex
}

func Periodic(fn func() error, d time.Duration) Check {
	return &periodic{fn: fn, d: d, chs: map[chan<- error]bool{}}
}

func (p *periodic) Notify(ch chan<- error) {
	p.m.Lock()
	defer p.m.Unlock()

	p.chs[ch] = true

	if len(p.chs) > 1 {
		return
	}

	p.done = make(chan struct{})
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

			err := p.fn()
			p.m.Lock()
			for ch := range p.chs {
				select {
				case ch <- err:
				default:
				}
			}
			p.m.Unlock()
		}
	}()
}

func (p *periodic) Stop(ch chan<- error) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.chs, ch)

	if len(p.chs) == 0 {
		close(p.done)
	}
}
