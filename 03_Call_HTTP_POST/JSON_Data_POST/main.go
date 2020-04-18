/* JSON data를 Post하는 것은 Plain Text를 Post하는 것과 비슷하다.
다만, data가 JSON format이므로 http.Post()의 두번째 parameter에 application/json을 적고,
세번째 parameter에 JSON으로 인코딩된 data를 전달하면 된다.
아래에서 JSON data는 encoding/json 표준 패키지의 Marshal() 함수를 써서,
임의의 구조체 data를 JSON으로 변경하는 방법을 사용했다.
*/
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Person data
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Eric", 28}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
