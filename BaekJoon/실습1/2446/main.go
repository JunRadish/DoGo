package main

import "fmt"

func main() {
	count := 0
	fmt.Scan(&count)
	for i := count; i >= 1; i-- {
		for j := count - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for j := 2*i - 1; j >= 1; j-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := 2; i <= count; i++ {
		for j := count - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
