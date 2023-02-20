package main

import "fmt"

func main() {

	//изменяемое число int64
	var n int64
	n = 55555555

	//i-й справа бит для изменения на 0 или 1
	i := 5

	//маска
	m := 1 << (i - 1)

	//изначчальное число
	fmt.Printf("%b\n", n)

	//измененное число
	fmt.Printf("%b\n", n^int64(m))

}
