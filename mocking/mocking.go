package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalCountdownOutput = "Go!"
const defaultCountdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

/* the following struct & method are for making a non-mocked version of function
where it is actually sleeping for a second, like:
	func main() {
		s := &ConfigurableSleeper{1 * time.Second, time.Sleep} // ... or this? ConfigurableSleeper{}
		Countdown(os.Stdout, s) // .................................. *s
	}
*/

func Countdown(out io.Writer, s Sleeper) {
	for i := defaultCountdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, finalCountdownOutput)
}
