package env

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	var cfg = struct {
		Duration time.Duration `env:"key=DURATION default=5s"`
	}{}

	if err := Process(&cfg); err != nil {
		t.Fatal(err)
	}

	if cfg.Duration != time.Second*5 {
		t.Fatalf("%v != %f ", cfg.Duration, time.Second*5)
	}
}
