package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	sum := 0
	n := 0
	for n < 5 {
		scanner.Scan()
		temp := strings.Split(scanner.Text(), "\n")
		for i := range temp {
			t, _ := strconv.Atoi(temp[i])
			if t < 40 {
				t = 40
			}
			sum += t
		}
		n++
	}
	fmt.Fprintf(w, "%d", sum/5)
}
