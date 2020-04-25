package main

import "fmt"

func main() {
	a, b, c, temp := 0, 0, 0, 0
	fmt.Scan(&a, &b, &c)
	mul := a * b * c
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range arr {
		temp = mul
		num := 0
		for temp >= 1 {
			if temp%10 == arr[i] {
				num++
			}
			temp /= 10
		}
		fmt.Println(num)
	}
}
