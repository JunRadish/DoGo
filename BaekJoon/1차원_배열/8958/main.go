package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	w := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < count; i++ {
		score := 0
		temp := 0
		scanner.Scan()
		input := scanner.Text()
		for j := range input {
			if string(input[j]) == "O" {
				temp++
				score = score + temp
			} else {
				temp = 0
			}
		}
		fmt.Fprintf(w, "%d\n", score)
	}
	w.Flush()
}
