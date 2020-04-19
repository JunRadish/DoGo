package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	for i := count; i >= 1; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
	w.Flush()
}
