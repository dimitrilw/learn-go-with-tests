package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeats string", func(t *testing.T) {
		got := Repeat("x")
		want := "xxxxx"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("x"))
	// Output: xxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
