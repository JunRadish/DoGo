package main

import "fmt"

func main() {
	num := 0
	fmt.Scan(&num)
	temp(num)
}

func temp(a int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}
