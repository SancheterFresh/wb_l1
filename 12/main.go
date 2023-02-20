package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "three"}

	set := map[string]bool{}

	for _, v := range strings {
		set[v] = true
	}
	fmt.Println(set)
}
