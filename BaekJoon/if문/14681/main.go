package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	judge(a, b)
}
func judge(a, b int) {
	switch {
	case a > 0 && b > 0:
		fmt.Println(1)
	case a < 0 && b > 0:
		fmt.Println(2)
	case a > 0 && b < 0:
		fmt.Println(4)
	case a < 0 && b < 0:
		fmt.Println(3)
	}
}
