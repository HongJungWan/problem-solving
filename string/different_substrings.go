// TODO: 서로 다른 부분 문자열의 개수
// https://www.acmicpc.net/problem/11478

package main

import (
	"fmt"
)

func countUniqueSubstrings(s string) int {
	substrings := make(map[string]bool)
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			substrings[s[i:j]] = true
		}
	}
	return len(substrings)
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(countUniqueSubstrings(s))
}

// fmt.Println(countUniqueSubstrings("ababc"))
// fmt.Println(countUniqueSubstrings("xyzxyz"))

// 맵에서 키의 존재 여부만 중요할 때, Go 개발자들은 종종 불리언 타입을 사용
