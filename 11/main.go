package main

import "fmt"

func intersection(a, b []int) (c []int) {
	m := make(map[int]int)

	for _, v := range a {
		m[v] = 1
	}

	for _, v := range b {
		m[v]++
		if m[v] > 1 {
			c = append(c, v)
		}
		m[v]--
	}
	return c
}

func main() {
	a := []int{5, 1, 2, 7, 8}
	b := []int{2, 4, 4, 9, 8}
	c := intersection(a, b)
	fmt.Println(c)

}
