package main

import (
	"fmt"
	"strconv"
	"sync"
)

type myMap[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func newMyMap[K comparable, V any]() *myMap[K, V] {
	return &myMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *myMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m = make(map[K]V)
}

func (m *myMap[K, V]) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}

func (c *myMap[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *myMap[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *myMap[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func (c *myMap[K, V]) Keys() []K {
	c.mu.Lock()
	defer c.mu.Unlock()
	keys := make([]K, 0, len(c.m))
	for k := range c.m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	m := newMyMap[string, int]()
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		s := "№" + strconv.Itoa(i)
		go func(key string, val int) {
			defer wg.Done()
			m.Set(key, val)
		}(s, i)
	}

	wg.Wait()

	k := m.Keys()
	for _, key := range k {
		res, ok := m.Get(key)
		if ok {
			fmt.Printf("Key %s = %d\n", key, res)
		}
	}

}
