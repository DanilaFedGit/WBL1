package main

import (
	"fmt"
	"math"
)

/*структура Point*/
type Point struct {
	x float64
	y float64
}

/* конструктор Point*/
func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

/* метод для поиска расстояния между двумя точками*/
func (point *Point) Distance(point2 *Point) float64 {
	return math.Sqrt(math.Pow(point.x-point2.x, 2) + math.Pow(point.y-point2.y, 2))
}
func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(6, 6)
	fmt.Println(point1.Distance(point2))

}
