package health

type Option interface {
	Apply(*Monitor)
}

type OptionFunc func(*Monitor)

func (o OptionFunc) Apply(m *Monitor) {
	o(m)
}

func WithChecks(checks ...Check) Option {
	return OptionFunc(func(m *Monitor) {
		for _, c := range checks {
			m.services = append(m.services, service{check: c})
		}
	})
}
