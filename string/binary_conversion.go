// 문자열 s가 "1"이 될 때까지 계속해서 이진 변환을 가했을 때,
//   1. 이진 변환 횟수와
//   2. 변환 과정에서 제거된 모든 0의 개수
//   각각 배열에 담아 return
// https://school.programmers.co.kr/learn/courses/30/lessons/70129

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryConversion(s string) []int {
	answer := make([]int, 2)

	for s != "1" {
		zeroCount := strings.Count(s, "0")
		answer[1] += zeroCount

		s = strings.ReplaceAll(s, "0", "")
		s = strconv.FormatInt(int64(len(s)), 2)
		answer[0]++
	}
	return answer
}

func main() {
	fmt.Println(binaryConversion("110010101001")) // Output: [3, 8]
	fmt.Println(binaryConversion("01110"))        // Output: [3, 3]
	fmt.Println(binaryConversion("1111111"))      // Output: [4, 1]
}
