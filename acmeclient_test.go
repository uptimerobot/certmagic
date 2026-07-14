package certmagic

import (
	"testing"
	"time"
)

func TestGetenv(t *testing.T) {
	if v := getenv("CERTMAGIC_TEST_UNSET_VAR", "fallback"); v != "fallback" {
		t.Errorf("expected fallback, got %q", v)
	}
	t.Setenv("CERTMAGIC_TEST_SET_VAR", "value")
	if v := getenv("CERTMAGIC_TEST_SET_VAR", "fallback"); v != "value" {
		t.Errorf("expected value, got %q", v)
	}
}

// Regression test: RateLimitEventsWindow used to be computed as
// limitWindow * time.Minute, double-applying the time unit since
// limitWindow is already a parsed time.Duration (e.g. "1m" became ~41 days).
func TestRateLimitEventsWindowDefault(t *testing.T) {
	want := time.Minute
	if RateLimitEventsWindow != want {
		t.Errorf("RateLimitEventsWindow = %v, want %v", RateLimitEventsWindow, want)
	}
}
