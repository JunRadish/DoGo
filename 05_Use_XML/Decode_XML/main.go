package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	fi, err := os.Open("/mnt/d/Workspace/test_files/encodedxml.txt")
	if err != nil {
		panic(err) // 만약 파일 못찾는 경우 err 발생
	}
	defer fi.Close()
	readBytes := make([]byte, 1024)
	// num1, err := fi.Read(xmlBytes)

	var mem Member
	count, err := fi.Read(readBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, readBytes[:count])
	err = xml.Unmarshal(readBytes[:count], &mem)
	// err = xml.Unmarshal(xmlBytes, &mem)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
