package health

import "sync"

type service struct {
	check  Check
	ch     chan error
	status Status
	err    error
	m      sync.RWMutex
}

func (s *service) loop(done <-chan struct{}) {
	ch := make(chan error)
	s.check.Notify(ch)

	s.m.Lock()
	s.status = Starting
	s.m.Unlock()

L:
	for {
		select {
		case <-done:
			s.check.Stop(ch)
			break L

		case err := <-ch:
			s.m.Lock()
			s.err = err
			if err != nil {
				s.status = Unhealthy
			} else {
				s.status = Healthy
			}
			s.m.Unlock()
		}
	}

	s.m.Lock()
	s.status = Stopped
	s.m.Unlock()
}
