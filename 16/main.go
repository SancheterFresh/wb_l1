package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	l := 0
	r := len(arr) - 1

	pivot := len(arr) / 2

	arr[pivot], arr[r] = arr[r], arr[pivot]

	for i := range arr {
		if arr[i] < arr[r] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]

	quicksort(arr[:l])
	quicksort(arr[l+1:])

	return arr
}

func main() {
	unsorted := []int{3, 6, 8, 4, 2, 4, 6, 241, 5, 4, 56, 133}
	fmt.Println(unsorted)
	fmt.Println(quicksort(unsorted))
}
