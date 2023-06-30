package template

import (
	"math"
)

type Void struct{}

// =============================================================================
// Min

const MAX = math.MaxInt

func Min(N ...int) (r int) {
	r = MAX
	for _, n := range N {
		if n < r {
			r = n
		}
	}
	return
}

// =============================================================================
// Max

const MIN = math.MinInt

func Max(N ...int) (r int) {
	r = MIN
	for _, n := range N {
		if n > r {
			r = n
		}
	}
	return
}
