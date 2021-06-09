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
		m.checks = append(m.checks, checks...)
	})
}

func OnFailure(fn func()) Option {
	return OptionFunc(func(m *Monitor) {
		m.failfns = append(m.failfns, fn)
	})
}
