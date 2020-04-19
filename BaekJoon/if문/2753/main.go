package main

import "fmt"

func lunar(a int) {
	if a%4 == 0 {
		if a%100 != 0 || a%400 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else {
		if a%400 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func main() {
	lun := 0
	fmt.Scanln(&lun)
	lunar(lun)
}
