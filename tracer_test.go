package tracer

import (
	"math/rand"
	"testing"
	"time"
)

func TestStatsd(t *testing.T) {
	for i := 0; i < 500; i++ {
		Statsd("UnitTests.Tracer", time.Now().Add(-1*time.Duration(rand.Intn(5))*time.Second))
	}
}
