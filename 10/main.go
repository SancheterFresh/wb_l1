package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	grouped := make(map[int][]float64)

	for _, v := range temperatures {
		key := int(v) / 10
		grouped[key] = append(grouped[key], v)

	}
	fmt.Println(grouped)

}
