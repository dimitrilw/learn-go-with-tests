package template

import (
	"fmt"
	"math"
	"testing"
)

// func assertEqual(t testing.TB, got, want interface{}) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestMin(t *testing.T) {
	t.Run("No arg", func(t *testing.T) {
		got := Min()
		want := math.MaxInt
		assertEqual(t, got, want)
	})
	t.Run("CSV args", func(t *testing.T) {
		got := Min(1, 2, 3)
		want := 1
		assertEqual(t, got, want)
	})
	t.Run("Slice expansion args", func(t *testing.T) {
		// cSpell:ignore nums
		nums := []int{1, 2, 3}
		got := Min(nums...)
		want := 1
		assertEqual(t, got, want)
	})
}

func TestMax(t *testing.T) {
	t.Run("No arg", func(t *testing.T) {
		got := Max()
		want := math.MinInt
		assertEqual(t, got, want)
	})
	t.Run("CSV args", func(t *testing.T) {
		got := Max(1, 2, 3)
		want := 3
		assertEqual(t, got, want)
	})
	t.Run("Slice expansion args", func(t *testing.T) {
		// cSpell:ignore nums
		nums := []int{1, 2, 3}
		got := Max(nums...)
		want := 3
		assertEqual(t, got, want)
	})
}

// =============================================================================
// Examples

func ExampleMin() {
	fmt.Println(Min(1, 2, 3))
	// Output: 1
}

func ExampleMax() {
	fmt.Println(Max(1, 2, 3))
	// Output: 3
}

// =============================================================================
// Benchmarks

func BenchmarkMin(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		Min(nums...)
	}
}

func BenchmarkMax(b *testing.B) {
	nums := []int{}
	for i := 0; i < b.N; i++ {
		nums = append(nums, i)
		Max(nums...)
	}
}
