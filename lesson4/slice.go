package main

import (
	"fmt"
)

func main() {
	arr := [100]int{0}
	slice := arr[:]
	for i := 0; i < 100; i++ {
		slice[i] = i + 1
	}

	slice = append(slice[:10], slice[90:]...)
	for i := 0; i < len(slice)/2; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-i-1]
		slice[len(slice)-i-1] = temp
	}

	fmt.Println(slice) // [1 2 3 4 5 6 7] 7 7
}
