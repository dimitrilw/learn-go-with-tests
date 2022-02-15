package iteration

func Repeat(s string, n int) (repeat_string string) {
	for i := 0; i < n; i++ {
		repeat_string += s
	}
	return
}
