package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h Human) sayMyName() {
	fmt.Printf("My name is: " + h.name)
}

func (h Human) sayHello() {
	fmt.Printf("Hello! ")
}

func main() {
	h := Human{"Ivan"}

	a := Action{h}

	a.sayHello()
	a.sayMyName()

	act := Action{}
	act.sayHello()
	act.name = "Nikola"
	act.sayMyName()

}
