package main

import "fmt"

func main() {
	a := 0
	fmt.Scanln(&a)
	grade(a)
}
func grade(g int) {
	switch {
	case g >= 90:
		fmt.Println("A")
	case g >= 80:
		fmt.Println("B")
	case g >= 70:
		fmt.Println("C")
	case g >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
