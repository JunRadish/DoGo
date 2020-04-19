package main

import "fmt"

func main() {
	menu1, menu2 := 2000, 2000
	temp := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&temp)
		if i < 3 {
			if temp < menu1 {
				menu1 = temp
			}
		} else {
			if temp < menu2 {
				menu2 = temp
			}
		}
	}
	fmt.Print(menu1 + menu2 - 50)
}
