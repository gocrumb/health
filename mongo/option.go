package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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

func ReadPref(rp *readpref.ReadPref) Option {
	return OptionFunc(func(c *check) {
		c.rp = rp
	})
}
