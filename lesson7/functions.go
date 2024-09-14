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
		return func(x float64) float64 {
			return x * x * math.Sqrt(3.) / 4.0
		}, true
	case circle:
		return func(r float64) float64 {
			return math.Pi * r * r
		}, true
	case square:
		return func(a float64) float64 {
			return a * a
		}, true
	default:
		return nil, false
	}

}

func main() {
	var myFigure figures = 1
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(1.)
	fmt.Println(myArea)
}
