package main

import "fmt"

func main() {
	count := 0
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := count - 1; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
