package syncthreads

import "sync"

type Counter struct {
	lock  sync.Mutex
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
}

func NewCounter() *Counter {
	return &Counter{}
}
