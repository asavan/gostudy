package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	p := Person{Name: "Aлекс", Email: "alex@yandex.ru"}
	res, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(res))
	} else {
		fmt.Println(err)
	}
}
