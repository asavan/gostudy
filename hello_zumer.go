package main

import "fmt"

func hi(year int) {
	var name string
	switch {
	case year >= 1946 && year <= 1964:
		name = "бумер"
	case year <= 1980:
		name = "представитель X"
	case year <= 1996:
		name = "миллениал"
	case year <= 2012:
		name = "зумер"
	default:
		name = "альфа"
	}
	if name != "" {
		fmt.Println("«Привет,", name+"!»")
	}
}

func main() {
	var i int
	fmt.Scan(&i)
	hi(i)
}
