/*Request Object를 생성하여 셋팅을 변경한 후
http.Client를 통해 호출할 수도 있다.*/

package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Person Data
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Eric", 28}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	// Request Object Creation
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	// Content-Type Header 추가
	req.Header.Add("Content-Type", "application/xml")

	// Client object에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response Check.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
