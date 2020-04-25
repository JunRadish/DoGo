package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z"}
	scanner.Scan()
	text := scanner.Text()

	for i := range alpha {
		flag := true
		for j := range text {
			if alpha[i] == string(text[j]) {
				fmt.Fprintf(w, "%d ", j)
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintf(w, "%d ", -1)
		}
	}
	w.Flush()
}
