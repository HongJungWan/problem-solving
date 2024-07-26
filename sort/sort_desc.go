// TODO: 문자열 내림차순으로 배치하기
// https://school.programmers.co.kr/learn/courses/30/lessons/12917

package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortDesc(str string) string {
	chars := strings.Split(str, "")

	// 사용자 정의 정렬 함수
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] > chars[j]
	})

	return strings.Join(chars, "")
}

func main() {
	s := "Zbcdefg"
	fmt.Println(sortDesc(s))
}
