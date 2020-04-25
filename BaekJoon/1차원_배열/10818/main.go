/*NewScanner 를 이용해서 bufio.Scanner 를 만들면 내부 버퍼 크기는 65536 byte 입니다.

Scan() 한 번에 읽어와야 될 데이터 (초기값은 '한 줄') 가 버퍼 크기를 넘어가면 제대로 읽어올수 없습니다.



NewScanner 직후, 다음 중 한 가지 방법을 적용하세요

1) scanner.Buffer 로 한 번에 읽을 버퍼 크기를 조절한다

2) scanner.Split 로 Scan 의 단위를 조절한다. 예를 들어 scanner.Split(bufio.ScanWords) 를 사용하면 단어 하나 단위로 Scan 할 수 있습니다.



이 문제에서는 2번 방법을 추천합니다.*/
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
	min, max := 1000001, -1000001
	scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Print(scanner.Text())
	// }
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < count; i++ {
		scanner.Scan()

		input := scanner.Text()
		target, _ := strconv.Atoi(input)
		if max < target {
			max = target
		}
		if min > target {
			min = target
		}
	}
	fmt.Fprintf(w, "%d %d", min, max)
	w.Flush()

}
