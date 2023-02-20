package main

import (
	"math/rand"
	"time"
)

var justString string

const symbols = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"

func createHugeString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

func someFunc() {
	v := createHugeString(1 << 10)
	b := make([]byte, 100)
	for i := 0; i < 100; i++ {
		b[i] = v[i]
	}

	justString = string(b)
}

func main() {
	someFunc()
	println(justString)
}
