package main

import (
	"fmt"
)

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int)
	result := make([]string, 0)
	for _, v := range input {
		_, ok := m[v]
		if !ok {
			result = append(result, v)
		}
		m[v] = 1
	}
	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	result := RemoveDuplicates(input)
	fmt.Println(result)
}
