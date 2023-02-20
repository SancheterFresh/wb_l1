package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func NewPoint(i, j float64) *point {
	return &point{
		x: i,
		y: j,
	}
}

func distBetween(p1, p2 *point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := NewPoint(4, 6)
	p2 := NewPoint(2, 10)
	fmt.Println(distBetween(p1, p2))

}
