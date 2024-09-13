package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(id figures) (func(float64) float64, bool) {
	switch id {
	case triangle:
		return func(f float64) float64 {
			return f * f * math.Sqrt(3.) / 2.0
		}, true
	case circle:
		return func(r float64) float64 {
			return math.Pi * r * r
		}, true
	case square:
		return func(f float64) float64 {
			return f * f
		}, true
	default:
		return func(f float64) float64 {
			return f
		}, false
	}

}

func main() {
	myFigure := triangle
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(1.)
	fmt.Println(myArea)
}
