package health

import "log/slog"

type Option interface {
	Apply(*Monitor)
}

type OptionFunc func(*Monitor)

func (o OptionFunc) Apply(m *Monitor) {
	o(m)
}

func Checks(checks ...Check) Option {
	return OptionFunc(func(m *Monitor) {
		m.checks = append(m.checks, checks...)
	})
}

func Logger(l *slog.Logger) Option {
	return OptionFunc(func(m *Monitor) {
		m.l = l
	})
}

func OnFailure(fn func()) Option {
	return OptionFunc(func(m *Monitor) {
		m.failfns = append(m.failfns, fn)
	})
}
