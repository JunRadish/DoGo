/*XML data를 Post하는 것은 JSON과 비슷하다.
encoding/xml package로 마샬링하고 MIME타입을 application/xml로 지정함.*/
package main

import (
	"bytes"
	"encoding/xml"
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
	pbytes, _ := xml.Marshal(person) // xml marshal
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/xml", buff)
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
