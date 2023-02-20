package main

import "fmt"

func square(value int, c *chan int) {
	*c <- value * value
}

func main() {
	c := make(chan int)
	nums := []int{2, 4, 6, 8, 10}

	for _, v := range nums {
		go square(v, &c)
		fmt.Println(<-c)
	}

}
