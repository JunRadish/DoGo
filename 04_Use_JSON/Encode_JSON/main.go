/*JSON format으로 encoding하기 위해서는 encoding/json 패키지의 Marshal()함수를 사용.
Go 구조체 혹은 map data를 JSON으로 Encoding하게 되는데,
해당 Go data value를 json.Marshal()의 parameter로 전달하면,
JSON으로 인코딩된 바이트배열과 에러객체를 리턴한다.
유의사항 : JSON의 Key는 문자열.
struct는 자동으로 field명을 string으로 사용.
map인 경우 map[string]T 처럼 Key가 String인 map만 지원.
*/

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

	//Go data
	mem := Member{"Eric", 28, true}

	//JSON encoding
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	//JSON byte --> string 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}
