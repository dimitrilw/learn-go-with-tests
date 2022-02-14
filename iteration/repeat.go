package iteration

func Repeat(s string) (repeat_string string) {
	for i := 0; i < 5; i++ {
		repeat_string += s
	}
	return
}
