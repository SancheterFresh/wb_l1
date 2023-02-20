package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	s = strings.ToLower(s)
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				return false

			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("abCdefAf"))
	fmt.Println(checkUnique("aabcd"))
	fmt.Println(checkUnique("ЯщлЩдЦц"))
	fmt.Println(checkUnique("ящлдщ"))
	fmt.Println(checkUnique("ящлдцук"))

}
