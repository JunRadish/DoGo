package main

import "fmt"

func main() {
	n, count := 100, 0
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		digit := 1
		temp := i
		for {
			if temp < 10 {
				break
			}
			temp /= 10
			digit++
		}
		temp = i
		if digit >= 3 {
			x := make([]int, digit-1)
			for j := range x {
				x[j] = temp%10 - (temp/10)%10
				temp /= 10
			}
			temp = 0
			for j := range x {
				for k := j + 1; k < digit-1; k++ {
					if x[j] != x[k] {
						temp++
					}
				}
			}
			if temp == 0 {
				count++
			}
		} else {
			count++
		}
		// 한수 판정
	}
	fmt.Print(count)
}
