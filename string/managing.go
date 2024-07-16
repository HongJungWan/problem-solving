// TODO: 문자열 다루기 기본
// https://school.programmers.co.kr/learn/courses/30/lessons/12918?language=go

package main

import (
	"fmt"
	"unicode"
)

func stringManaging(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}

	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(stringManaging("a234"))
	fmt.Println(stringManaging("1234"))
}
