package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	// quickCheckIdentity := func(a int) bool {
	// 	left := Add(a, 0)
	// 	right := Add(0, a)
	// 	return left == a && right == a && left == right
	// }
	// quickCheckAssociativity := func(a, b, c int) bool {
	// 	return Add(a, Add(b, c)) == Add(Add(a, b), c)
	// }

	t.Run("adds numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}

func ExampleAdd() {
	fmt.Println(Add(8, 9))
	// Output: 17
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i)
	}
}
