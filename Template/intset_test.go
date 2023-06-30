package template

import (
	"fmt"
	"sort"
	"testing"
)

/*
func (s IntSet) Has(n int) bool { _, ok := s[n]; return ok }
func (s *IntSet) Add(n int)     { (*s)[n] = Void{} }
func (s *IntSet) Del(n int) {
	if _, ok := (*s)[n]; ok {
		delete(*s, n)
	}
}
func (s IntSet) Len() int { return len(s) }
func (s IntSet) Slice() (r []int) {
	r = make([]int, len(s))
	i := 0
	for n := range s {
		r[i] = n
		i++
	}
	return
}
//*/

func TestIntSet(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		tSet := IntSet{}
		assertEqual(t, len(tSet), 0)

		tSet.Add(1)
		tSet.Add(2)
		tSet.Add(3)
		assertEqual(t, len(tSet), 3)
	})
	t.Run("Has", func(t *testing.T) {
		tSet := IntSet{}
		tSet.Add(1)
		assertEqual(t, tSet.Has(1), true)
		assertEqual(t, tSet.Has(2), false)
	})
	t.Run("Del", func(t *testing.T) {
		tSet := IntSet{}
		tSet.Add(1)
		tSet.Add(2)
		tSet.Add(3)
		assertEqual(t, len(tSet), 3)
		tSet.Del(2)
		assertEqual(t, len(tSet), 2)
		assertEqual(t, tSet.Has(2), false)
	})
	t.Run("Slice", func(t *testing.T) {
		tSet := IntSet{}
		tSet.Add(1)
		tSet.Add(2)
		tSet.Add(3)

		s := tSet.Slice()
		sort.Ints(s)

		want := []int{1, 2, 3}
		for i, got := range tSet.Slice() {
			assertEqual(t, got, want[i])
		}
	})
}

// =============================================================================
// Examples

func ExampleIntSet() {
	fmt.Println(Min(1, 2, 3))
	// Output: 1
}

// =============================================================================
// Benchmarks

func BenchmarkIntSet(b *testing.B) {
	// setup
	tSet := IntSet{}
	for i := 0; i < b.N; i++ {
		tSet.Add(i)
	}
	b.ResetTimer() // needed??

	b.Run("Slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tSet.Slice()
		}
	})
}
