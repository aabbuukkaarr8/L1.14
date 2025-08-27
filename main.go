package main

import (
	"fmt"
)

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("Тип: int, значение =", val)
	case string:
		fmt.Println("Тип: string, значение =", val)
	case bool:
		fmt.Println("Тип: bool, значение =", val)
	case chan int:
		fmt.Println("Тип: chan int")
	case chan string:
		fmt.Println("Тип: chan string")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
	detectType(make(chan string))
}
