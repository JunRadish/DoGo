package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < count; i++ {
		sum := 0
		var overAve float32 = 0
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())
		x := make([]int, input)
		for j := 0; j < input; j++ {
			scanner.Scan()
			x[j], _ = strconv.Atoi(scanner.Text())
			sum += x[j]
		}
		for j := 0; j < input; j++ {
			if (float32(sum) / float32(input)) < float32(x[j]) {
				overAve++
			}
		}
		fmt.Fprintf(w, "%.3f%%\n", (overAve/float32(input))*100)
	}
	w.Flush()
}
