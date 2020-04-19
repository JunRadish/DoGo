/*Go data를 XML으로 인코딩 하기위해서 encoding/xml package의 Marshal()을 사용.
Encoding한 byte arrayfmf string으로 변경해야하면
string(byte array)로 가능*/

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Member
type Member struct {
	Name   string
	Age    int
	Active bool
}
type Members struct {
	Member []Member
}

func main() {
	mem := Member{"Eric", 28, true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}
	xmlString := string(xmlBytes)
	fmt.Printf(xmlString)

	fo, err := os.Create("/mnt/d/Workspace/test_files/encodedxml.xml")
	if err != nil {
		panic(err) // 경로가 존재하지 않는 경우 err 발생
	}
	defer fo.Close()

	_, err = fo.Write(xmlBytes)
	if err != nil {
		panic(err)
	}

}
