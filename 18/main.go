package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c  int
	mx sync.Mutex
}

func NewCounter() *counter {
	return &counter{
		c:  0,
		mx: sync.Mutex{},
	}
}

func (cnt *counter) Increment(wg *sync.WaitGroup) {
	cnt.mx.Lock()
	cnt.c++
	cnt.mx.Unlock()
	wg.Done()
}

func main() {

	counter := *NewCounter()

	var wg sync.WaitGroup

	for i := 0; i < 55; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}
	wg.Wait()

	fmt.Println(counter.c)

}
