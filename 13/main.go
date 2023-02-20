package main

import "fmt"

func main() {
	a := 6
	b := 7

	fmt.Println(a, " ", b)

	//способ 1
	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, " ", b)

	//способ 2
	a, b = b, a

	fmt.Println(a, " ", b)

}
