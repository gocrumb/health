// Copyright 2021 Furqan Software Ltd. All rights reserved.

package mongo

import (
	"context"
	"time"

	"github.com/gocrumb/health"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DefaultPeriod = 5 * time.Minute
)

type check struct {
	c      *mongo.Client
	period time.Duration
	rp     *readpref.ReadPref
}

func Ping(c *mongo.Client, opts ...Option) health.Check {
	k := check{c: c, period: DefaultPeriod}
	for _, o := range opts {
		o.apply(&k)
	}
	return health.Periodic(k.run, k.period)
}

func (k check) run() error {
	return k.c.Ping(context.TODO(), k.rp)
}
