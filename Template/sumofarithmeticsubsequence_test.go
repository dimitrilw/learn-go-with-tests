package template

import (
	"fmt"
	"testing"
)

func TestSumofArithmeticSubSequence(t *testing.T) {
	t.Run("same number for args", func(t *testing.T) {
		got := SumOfArithmeticSubSequence(1, 1)
		want := 1
		assertEqual(t, got, want)

		got = SumOfArithmeticSubSequence(9, 9)
		want = 9
		assertEqual(t, got, want)
	})

	t.Run("different numbers for args", func(t *testing.T) {
		got := SumOfArithmeticSubSequence(1, 5)
		want := 15
		assertEqual(t, got, want)

		got = SumOfArithmeticSubSequence(3, 5)
		want = 12
		assertEqual(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleSumOfArithmeticSubSequence() {
	fmt.Println(SumOfArithmeticSubSequence(1, 5))
	// Output: 15
}

// =============================================================================
// Benchmarks

func BenchmarkSumOfArithmeticSubSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfArithmeticSubSequence(1, i)
	}
}
