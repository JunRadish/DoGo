package main

import "fmt"

func main() {
	max, input, temp := 0, 0, 0
	for i := 1; i <= 9; i++ {
		fmt.Scan(&input)
		if max < input {
			max = input
			temp = i
		}
	}
	fmt.Printf("%d\n%d", max, temp)
}
