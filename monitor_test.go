package health_test

import (
	"time"

	"github.com/grumbs/health"
	"github.com/grumbs/health/http"
)

func ExampleMonitor() {
	m := health.New(
		health.WithChecks(
			http.Head("http://example.com/", http.Period(1*time.Minute)),
		),
	)
	m.Start()
}
