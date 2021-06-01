package http

import "time"

type Option interface {
	apply(*check)
}

type OptionFunc func(*check)

func (o OptionFunc) apply(c *check) {
	o(c)
}

func Period(d time.Duration) Option {
	return OptionFunc(func(c *check) {
		c.period = d
	})
}
