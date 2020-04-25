package main

import "fmt"

func main() {
	arr := [10]int{}
	temp := 0
	for i := range arr {
		fmt.Scan(&arr[i])
		arr[i] %= 42
	}
	for i := range arr {
		for j := i + 1; j < 10; j++ {
			if arr[i] != arr[j] {
				temp++
				break
			}
		}
	}
	fmt.Print(temp)
}
