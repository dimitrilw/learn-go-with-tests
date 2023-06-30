package template

import (
	"testing"
)

func assertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
