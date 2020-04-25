package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scanln(&a, &b, &c)
	switch {
	case a <= b && b <= c || c <= b && b <= a:
		fmt.Print(b)
	case b <= a && a <= c || c <= a && a <= b:
		fmt.Print(a)
	case a <= c && c <= b || b <= c && c <= a:
		fmt.Print(c)
	}
}
