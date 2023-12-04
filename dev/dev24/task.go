package main

import (
	"L1/dev/dev24/point"
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func main() {
	p1 := point.NewPoint(19.23425, 4593.12412)
	p2 := point.NewPoint(2349423.1234, 346.124)
	x1, y1 := p1.GetCoordinates()
	x2, y2 := p2.GetCoordinates()
	fmt.Printf("Distance is %f", distance(x1, y1, x2, y2))
}
