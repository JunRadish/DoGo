package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*Go의 표준 패키지인 http 패키지는 웹 관련 클라이언트 및 서버 기능을 제공한다.
그 중 http.Get() 메서드는 쉽게 웹 페이지나 웹 데이타들 가져오는데 사용된다.*/
func main() {
	// Calling GET
	resp, err := http.Get("https://www.naver.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Result Printing
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
