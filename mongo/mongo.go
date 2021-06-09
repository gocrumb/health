// Copyright 2021 Furqan Software Ltd. All rights reserved.

package mongo

import (
	"time"

	"github.com/gocrumb/health"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DefaultPeriod = 5 * time.Minute
)

type check struct {
	c      *mongo.Client
	period time.Duration
}

func Ping(c *mongo.Client, opts ...Option) health.Check {
	k := check{c, DefaultPeriod}
	for _, o := range opts {
		o.apply(&k)
	}
	return health.Periodic(k.run, k.period)
}

func (k check) run() error {
	return k.c.Ping()
}
