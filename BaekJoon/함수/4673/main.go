package main

import "fmt"

func main() {
	// selfNum := 0
	num := 10000
	a := make([]int, num)
	for i := 1; i <= num; i++ {
		temp := i
		sum := i
		for temp/10 > 0 {
			sum += temp % 10
			temp /= 10
		}
		sum += temp
		a[i-1] = sum
	}
	for i := 1; i <= num; i++ {
		flag := false
		for j := 1; j <= num; j++ {
			if a[j-1] == i {
				flag = true
			}
		}
		if !flag {
			fmt.Println(i)
		}
		// fmt.Println(a[i-1])
	}
}
