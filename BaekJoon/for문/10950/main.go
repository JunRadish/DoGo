package main

import "fmt"

func main() {
	k, num1, num2 := 0, 0, 0
	fmt.Scan(&k)
	x := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanln(&num1, &num2)
		x[i] = num1 + num2
	}
	for i := range x {
		fmt.Println(x[i])
	}
}
