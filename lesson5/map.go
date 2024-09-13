package main

import (
	"fmt"
)

func findIndex(a []int, k int) []int {
	m := make(map[int]int)
	for i, v := range a {
		prev, ok := m[k-v]
		if ok {
			return []int{prev, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	m := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}
	for k, v := range m {
		if v > 500 {
			fmt.Println(k)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var result int
	for _, item := range order {
		result += m[item]
	}
	fmt.Println(result)
}
