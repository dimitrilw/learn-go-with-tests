package template

// type Void struct{}

type IntSet map[int]Void

func (s IntSet) Has(n int) bool { _, ok := s[n]; return ok }
func (s *IntSet) Add(n int)     { (*s)[n] = Void{} }
func (s *IntSet) Del(n int) {
	if _, ok := (*s)[n]; ok {
		delete(*s, n)
	}
}

/*
Built using an incrementing index (i) because it is faster than append.

	Benchmark results:
		BenchmarkIntSet/SliceViaIncrement-10   31304653   37.05 ns/op   8 B/op   1 allocs/op
		BenchmarkIntSet/SliceViaAppend-10      28361586   42.09 ns/op   8 B/op   1 allocs/op
*/
func (s IntSet) Slice() (r []int) {
	r = make([]int, len(s))
	i := 0
	for k := range s {
		r[i] = k
		i++
	}
	return
}
