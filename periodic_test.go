package health_test

import (
	"testing"
	"time"

	"github.com/gocrumb/health"
)

func TestPeriodic(t *testing.T) {
	n := 0
	fn := func() error { n++; return nil }
	c := health.Periodic(fn, 250*time.Microsecond)
	ch := make(chan error)
	c.Notify(ch)
	<-ch
	<-ch
	<-ch
	if n != 3 {
		t.Fatalf("want n == 3, got %d", n)
	}
}
