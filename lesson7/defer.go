package main

import (
	"fmt"
)

var Global = 5

func changeValue() {
	defer func(prev int) {
		Global = prev
	}(Global)
	Global = 7
	fmt.Println(Global)
}

func main() {
	changeValue()
	fmt.Println(Global)
}
