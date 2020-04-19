package main

import "fmt"

func rem(a int, b int, c int) {
	fmt.Println((a + b) % c)
	fmt.Println(((a % c) + (b % c)) % c)
	fmt.Println((a * b) % c)
	fmt.Println(((a % c) * (b % c)) % c)
}

func main() {
	a, b, c := 1, 1, 1
	fmt.Scanln(&a, &b, &c)
	rem(a,b,c)
}
