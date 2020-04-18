/*http.PostForm()은 form data를 보내는데 유용한 method이다.
즉, 일반 웹페이지에서 Submit버튼을 누르면,
입력 컨트롤러들이 data를 Form Data로 서버에 전송하는데,
이 PostForm()함수를 사용하면 동일한 효과를 내며 data를 쉽게 webserver로 보낼 수 있다.
아래 예제는 Name과 Age data를 form 형식으로 webserver에 전송하는 예이다.
httpbin.org/post는 전송한 data를 그대로 return하므로 Response에서는
form 필드에 Name과 Age data가 동일함을 볼 수 있다.
*/
package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 간단한 http.PostForm 예제
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//Response Check.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
