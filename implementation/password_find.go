// TODO: 비밀번호 찾기
// https://www.acmicpc.net/problem/17219

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 사이트 주소와 비밀번호를 저장할 맵 생성
	passwords := make(map[string]string)

	// 사이트 주소와 비밀번호를 입력받아 맵에 저장
	for i := 0; i < N; i++ {
		var site, password string
		fmt.Fscanln(reader, &site, &password)
		passwords[site] = password
	}

	// 비밀번호를 찾으려는 사이트 주소를 입력받아 비밀번호 출력
	for i := 0; i < M; i++ {
		var site string
		fmt.Fscanln(reader, &site)
		fmt.Fprintln(writer, passwords[site])
	}
}
