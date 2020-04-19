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

	for {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(inputs[0])
		b, _ := strconv.Atoi(inputs[1])
		if a == 0 && b == 0 {
			break
		}
		fmt.Fprintf(w, "%d\n", a+b)
	}
	w.Flush()
}
