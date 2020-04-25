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
	max := 0
	var sum float32

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	x := make([]int, count)

	for i := 0; i < count; i++ {
		scanner.Scan()

		input := scanner.Text()
		target, _ := strconv.Atoi(input)
		x[i] = target

		if max < target {
			max = target
		}
	}
	for i := range x {
		sum += (float32(x[i]) / float32(max)) * float32(100)
		// fmt.Println(sum)
	}
	fmt.Print(sum / float32(count))
}
