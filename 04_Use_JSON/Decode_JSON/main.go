/*JSON으로 encoding된 data를 다시 decoding하기 위해서는 encoding/json package의
Unmarshal()함수를 사용.
Unmarshal()함수의 첫번째 parameter에는 JSON data,
두번째 parameter에는 출력할 struct or map을 포인터로 지정.
Return value는 Error 객체이고, Error가 없을 경우,
두번째 parameter에 원래 data가 복원된다.*/

package main

import (
	"encoding/json"
	"fmt"
)

//Member data
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// Test용 JSON data
	jsonBytes, _ := json.Marshal(Member{"Eric", 28, true})

	//JSON Decoding
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	// mem struct field access
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
