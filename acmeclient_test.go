package certmagic

import (
	"testing"
	"time"
)

func TestRateLimitEventsFromEnv(t *testing.T) {
	if n := rateLimitEventsFromEnv("CERTMAGIC_TEST_UNSET_EVENTS", 60); n != 60 {
		t.Errorf("expected fallback 60, got %d", n)
	}
	t.Setenv("CERTMAGIC_TEST_EVENTS", "25")
	if n := rateLimitEventsFromEnv("CERTMAGIC_TEST_EVENTS", 60); n != 25 {
		t.Errorf("expected 25, got %d", n)
	}
	for _, invalid := range []string{"0", "-5", "not-a-number", "999999999"} {
		t.Setenv("CERTMAGIC_TEST_EVENTS_INVALID", invalid)
		if n := rateLimitEventsFromEnv("CERTMAGIC_TEST_EVENTS_INVALID", 60); n != 60 {
			t.Errorf("value %q: expected fallback 60, got %d", invalid, n)
		}
	}
}

func TestRateLimitWindowFromEnv(t *testing.T) {
	if d := rateLimitWindowFromEnv("CERTMAGIC_TEST_UNSET_WINDOW", time.Minute); d != time.Minute {
		t.Errorf("expected fallback 1m, got %v", d)
	}
	t.Setenv("CERTMAGIC_TEST_WINDOW", "90s")
	if d := rateLimitWindowFromEnv("CERTMAGIC_TEST_WINDOW", time.Minute); d != 90*time.Second {
		t.Errorf("expected 90s, got %v", d)
	}
	for _, invalid := range []string{"not-a-duration", "500ms"} {
		t.Setenv("CERTMAGIC_TEST_WINDOW_INVALID", invalid)
		if d := rateLimitWindowFromEnv("CERTMAGIC_TEST_WINDOW_INVALID", time.Minute); d != time.Minute {
			t.Errorf("value %q: expected fallback 1m, got %v", invalid, d)
		}
	}
}

// Regression test: RateLimitEventsWindow used to be computed as
// limitWindow * time.Minute, double-applying the time unit since
// limitWindow was already a parsed time.Duration (e.g. "1m" became ~41 days).
func TestRateLimitEventsWindowDefault(t *testing.T) {
	want := time.Minute
	if RateLimitEventsWindow != want {
		t.Errorf("RateLimitEventsWindow = %v, want %v", RateLimitEventsWindow, want)
	}
}
