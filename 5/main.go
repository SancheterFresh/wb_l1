package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sub(c *chan int) {
	for v := range *c {
		fmt.Println(v)
	}
}

func main() {
	c := make(chan int)

	fmt.Printf("Enter duration time: ")
	var n int
	fmt.Scan(&n)

	go sub(&c)

	rand.Seed(time.Now().UnixNano())

	timeout := time.After(time.Duration(n * int(time.Second)))

	to := true

	for to {
		select {
		case <-timeout:
			fmt.Printf("Time Out!\n")
			close(c)
			to = false
		default:
			c <- rand.Int()
		}

	}

}
