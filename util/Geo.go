package util

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Line struct {
	Points []Point
}

func (this *Line) print() {
	for _, p := range this.Points {
		fmt.Printf("(%f, %f) ", p.X, p.Y)
	}
	println("")
}

func PrintLines(result []Line) {
	println("___________________________________________________________________________________________________________________")
	println("Lines")
	println("___________________________________________________________________________________________________________________")
	for _, l := range result {
		l.print()
	}
	println("___________________________________________________________________________________________________________________")
}

func PointDist(a Point, b Point) float64 {
	dxsq := math.Pow(a.X-b.X, 2)
	dysq := math.Pow(a.Y-b.Y, 2)
	return math.Sqrt(dxsq + dysq)
}
