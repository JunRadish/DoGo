package main

import (
	"fmt"
)

func mul(a, b int) {
	temp := b
	for {
		if temp < 1 {
			break
		}
		fmt.Println(a * (temp % 10))
		temp /= 10
	}
	fmt.Println(a * b)
}
func main() {
	a, b := 472, 385
	fmt.Scan(&a, &b)
	mul(a, b)
}
