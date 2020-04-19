package main

import (
	"fmt"
)

func cal(a int, b int) {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	cal(a, b)
}
