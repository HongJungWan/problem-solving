// TODO: 접미사 배열
// https://www.acmicpc.net/problem/11656

package main

import (
	"fmt"
	"sort"
)

func sortedSuffixes(s string) {
	// 모든 접미사를 저장할 슬라이스 생성
	suffixes := make([]string, len(s))

	// 각 인덱스에서 시작하는 접미사를 슬라이스에 추가
	for i := 0; i < len(s); i++ {
		suffixes[i] = s[i:]
	}

	// 접미사들을 사전 순으로 오름차순 정렬
	sort.Strings(suffixes)

	// 정렬된 접미사들을 출력
	for _, suffix := range suffixes {
		fmt.Println(suffix)
	}
}

func main() {
	var s string
	fmt.Scan(&s)

	sortedSuffixes(s)
}
