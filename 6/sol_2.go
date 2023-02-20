package main

//завершение при закрытии канала

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range c {
			fmt.Println(v)
		}
		return
	}()

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	wg.Wait()

}
