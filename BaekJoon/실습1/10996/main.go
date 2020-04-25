package main

import "fmt"

func main() {
	count := 0
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			if j%2 != 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
		for j := 1; j <= count; j++ {
			if j%2 == 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}
