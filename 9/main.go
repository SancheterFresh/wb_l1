package main

import (
	"fmt"
)

func send(values []int, c *chan int) {
	for _, v := range values {
		*c <- v
	}
	close(*c)
}

func square(up *chan int, down *chan int) {

	for v := range *up {
		*down <- v * 2
	}
	close(*down)

}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	go send(nums, &c1)
	go square(&c1, &c2)

	for v := range c2 {
		fmt.Println(v)
	}

}
