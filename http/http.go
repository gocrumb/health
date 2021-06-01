package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/grumbs/health"
)

const (
	DefaultPeriod  = 5 * time.Minute
	DefaultTimeout = 1 * time.Minute
)

var ErrStatusCode = errors.New("got non-2xx status code")

type check struct {
	method  string
	url     string
	period  time.Duration
	timeout time.Duration
	done    chan struct{}
}

func Head(url string, opts ...Option) health.Check {
	c := check{"HEAD", url, DefaultPeriod, DefaultTimeout, make(chan struct{})}
	for _, o := range opts {
		o.apply(&c)
	}
	return health.Periodic(c.run, c.period)
}

func (c check) run() error {
	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return ErrStatusCode
	}
	return nil
}
