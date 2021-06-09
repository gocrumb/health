package health_test

import (
	"time"

	"github.com/gocrumb/health"
	"github.com/gocrumb/health/http"
)

func ExampleMonitor() {
	m := health.New(
		health.Checks(
			http.Head("http://example.com/", http.Period(1*time.Minute)),
		),
	)
	m.Start()
}
