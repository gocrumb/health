package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/grumbs/health"
)

const (
	DefaultPeriod = 5 * time.Minute
)

type check struct {
	c      *redis.Client
	period time.Duration
}

func Ping(c *redis.Client, opts ...Option) health.Check {
	k := check{c, DefaultPeriod}
	for _, o := range opts {
		o.apply(&k)
	}
	return health.Periodic(k.run, k.period)
}

func (k check) run() error {
	return k.c.Ping().Err()
}
