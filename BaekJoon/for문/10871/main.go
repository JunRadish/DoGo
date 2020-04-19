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

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	x, _ := strconv.Atoi(inputs[1])

	scanner.Scan()
	newInputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		temp, _ := strconv.Atoi(newInputs[i])
		if x > temp {
			fmt.Fprintf(w, "%d ", temp)
		}
	}
	// scanner.Scan()
	// count, _ := strconv.Atoi(scanner.Text())
	// for i := 1; i <= count; i++ {
	// 	scanner.Scan()
	// 	inputs := strings.Split(scanner.Text(), " ")
	// 	a, _ := strconv.Atoi(inputs[0])
	// 	b, _ := strconv.Atoi(inputs[1])
	// 	fmt.Fprintf(w, "Case #%d: %d + %d = %d\n", i, a,b,a+b)
	// }
	w.Flush()
}
