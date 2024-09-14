package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch v2 := a.(type) {
	case int:
		return v2 * b
	case string:
		var buf strings.Builder
		for i := 0; i < b; i++ {
			buf.WriteString(v2)
		}
		return buf.String()
	case fmt.Stringer:
		return strings.Repeat(v2.String(), b)
	default:
		return nil
	}
}

func main() {
	fmt.Println(Mul(1, 5))
	fmt.Println(Mul("a", 5))
}
