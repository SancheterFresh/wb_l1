package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := 5

	if i < 0 || i > len(a) {
		fmt.Println("Не соответствует размеру")
	} else {
		fmt.Println(a)
		a = append(a[0:i-1], a[i:]...)
		fmt.Println(a)

	}

}
