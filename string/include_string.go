// TODO: 문자열 집합
// https://www.acmicpc.net/problem/14425

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 지정된 수만큼 문자열을 읽어서 map으로 반환
func readStrings(count int, reader *bufio.Reader) map[string]bool {
	stringsSet := make(map[string]bool)

	for i := 0; i < count; i++ {
		text, err := reader.ReadString('\n')
		if err != nil && text == "" { // 더 이상 읽을 문자열이 없는 경우
			break
		}
		text = strings.TrimSpace(text) // 개행 문자를 제거
		stringsSet[text] = true
	}
	return stringsSet
}

// M개의 문자열 중 집합 S에 포함되어 있는 문자열의 수 찾기
func countIncludedString(M int, reader *bufio.Reader, setS map[string]bool) int {
	count := 0

	for i := 0; i < M; i++ {
		text, err := reader.ReadString('\n')
		if err != nil && text == "" {
			break
		}
		text = strings.TrimSpace(text)
		if setS[text] { // 집합 S에 포함되어 있는지 확인
			count++
		}
	}
	return count
}

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M) // 개행 문자를 명시적으로 처리

	reader := bufio.NewReader(os.Stdin)

	// 집합 S에 포함된 문자열 저장
	setS := readStrings(N, reader)
	count := countIncludedString(M, reader, setS)
	fmt.Println(count)
}
