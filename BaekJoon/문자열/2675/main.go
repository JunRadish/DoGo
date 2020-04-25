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
	var caseNum int
	fmt.Scan(&caseNum)
	for i := 0; i < caseNum; i++ {
		scanner.Scan()
		inputs, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		text := scanner.Text()
		for j := range text {
			for k := 0; k < inputs; k++ {
				fmt.Fprintf(w, "%s", string(text[j]))
			}
		}
		if i < caseNum-1 {
			fmt.Fprintf(w, "\n")
		}
	}
	w.Flush()
}
