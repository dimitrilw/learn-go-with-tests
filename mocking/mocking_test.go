package mocking

import (
	"bytes"
	"strings"
	"testing"

	"github.com/lithammer/dedent"
)

func TestCountdown(t *testing.T) {
	t.Run("with mocked Sleep", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpySleeper{}
		Countdown(b, s)

		got := b.String()
		want := strings.TrimSpace(dedent.Dedent(`
			3
			2
			1
			Go!
		`))

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		numCalls := 3
		if s.Calls != numCalls {
			t.Errorf("incorrect number of Sleep calls, got %d want %d", s.Calls, numCalls)
		}
	})
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
