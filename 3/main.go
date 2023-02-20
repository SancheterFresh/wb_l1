package main

import (
	"fmt"
)

func square(values []int, c *chan int) {
	for _, v := range values {
		*c <- v * v
	}
	close(*c)
}

func sum(c *chan int) {
	var sum int
	for v := range *c {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	c := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	go square(nums, &c)
	sum(&c)

}
