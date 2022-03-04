package mocking

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/lithammer/dedent"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3... to Go!", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpyCountdownOperations{}
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
	})

	t.Run("sleep between every write", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		lastCall := ""
		for _, call := range spySleepPrinter.Calls {
			if call == lastCall {
				t.Fatal(fmt.Sprintf("incorrect call order, should alternate write/sleep/write/etc, got %q", spySleepPrinter.Calls))
			}
			lastCall = call
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
