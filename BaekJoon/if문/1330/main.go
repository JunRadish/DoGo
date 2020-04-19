package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	comp(a, b)
}

func comp(a int, b int) {
	if a < b {
		fmt.Println("<")
	}
	if a > b {
		fmt.Println(">")
	}
	if a == b {
		fmt.Println("==")
	}
}
