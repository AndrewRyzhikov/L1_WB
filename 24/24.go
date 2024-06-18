package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func calculateDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

func main() {
	p1, p2 := NewPoint(2, 3), NewPoint(5, 0)
	fmt.Println(calculateDistance(*p1, *p2))
}
