package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	alarm(a, b)
}
func alarm(a, b int) {
	if a > 0 {
		if b < 45 {
			fmt.Println(a-1, b-45+60)
		} else {
			fmt.Println(a, b-45)
		}
	} else {
		if b < 45 {
			fmt.Println(a-1+24, b-45+60)
		} else {
			fmt.Println(a, b-45)
		}
	}
}
