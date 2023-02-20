package main

import (
	"fmt"
	"sort"
)

func binSearch(a []int, s int) bool {
	l := 0
	h := len(a) - 1

	for l <= h {
		m := (l + h) / 2
		if a[m] < s {
			l = m + 1
		} else if a[m] > s {
			h = m - 1
		} else if a[m] == s {
			return true
		}

	}
	return false

}

func main() {
	array := []int{5, 3, 7, 4, 434, 6, 76, 5, 4, 45, 7, 11, 67, 89, 45, 56, 95, 25, 85, 451, 1536, 6852, 72, 98, 7, 72}
	sort.Ints(array)

	fmt.Printf("Число для поиска: ")
	var n int
	fmt.Scan(&n)

	if binSearch(array, n) {
		fmt.Printf("%v найдено!", n)
	} else {
		fmt.Printf("%v не найдено!", n)
	}

}
