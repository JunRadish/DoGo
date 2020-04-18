/*Go의 Standard package인 http package는 웹 관련 클라이언트 및 서버 기능을 제공.package main
그 중 http.Post() method는 웹서버로 간단히 데이터를 POST 할 때 사용된다.
아래는 테스트 웹사이트인 httpbin.org에 임의의 텍스트를 POST를 사용해서 보내는 코드이다.
Post()메서드의 첫번째 파라미터는 Post를 받는 URL을 적고,
두번째는 Request Body의 MIME 타입을, 그리고 마지막에는 전송할 데이타를 (io.Reader로) 보낸다.
httpbin.org/post 는 전송한 데이타를 그대로 리턴하는데,
Response의 data를 체크하면 전송 데이타가 그대로 리턴됨을 볼 수 있다.
*/
package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	// 간단한 http.Post 예제
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
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
