package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x interface{}

	x = make(chan int)

	switch reflect.ValueOf(x).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("String")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Float64:
		fmt.Println("float")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Println("unknown")
	}
}
