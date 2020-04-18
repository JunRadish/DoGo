package main

import "io/ioutil"

func main(){
	// 파일 읽기
	bytes, err := ioutil.ReadFile("/mnt/d/Workspace/test_files/test.txt")
	if err != nil{
		panic(err)
	}
	// 파일 쓰기
	err = ioutil.WriteFile("/mnt/d/Workspace/test_files/output2.txt", bytes, 0)
	if err != nil{
		panic(err)
	}
}