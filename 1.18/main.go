package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m sync.Mutex
	n int
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.n += 1
}

func (c *Counter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func main() {
	var count Counter
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(count.Value())
}
