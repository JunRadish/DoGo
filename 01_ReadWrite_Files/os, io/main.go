package main

import (
	"io"
	"os"
)

func main() {
	// 입력파일 열기
	// WSL에서 구동하면, Linux이므로 file path 입력시 linux 형태로~
	fi, err := os.Open("/mnt/d/Workspace/test_files/test.txt")
	if err != nil {
		panic(err) // 만약 파일 못찾는 경우 err 발생
	}
	defer fi.Close()

	// 출력파일 생성
	fo, err := os.Create("/mnt/d/Workspace/test_files/output.txt")
	if err != nil {
		panic(err) // 경로가 존재하지 않는 경우 err 발생
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	// 루프
	for {
		// 읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}
		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
