package mocking

import (
	"fmt"
	"io"
)

const finalCountdownOutput = "Go!"
const defaultCountdownStart = 3

type Sleeper interface {
	Sleep()
}

/* the following struct & method are for making a non-mocked version of function
where it is actually sleeping for a second, like:
	func main() {
		s := &DefaultSleeper{}
		Countdown(os.Stdout, s)
	}
*/
// type DefaultSleeper struct {}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

func Countdown(out io.Writer, s Sleeper) {
	for i := defaultCountdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, finalCountdownOutput)
}
