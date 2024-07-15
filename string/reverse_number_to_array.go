// TODO: 자연수 n을 뒤집어 각 자리 숫자를 원소로 가지는 배열 형태로 리턴
// https://school.programmers.co.kr/learn/courses/30/lessons/12932

package main

import (
	"fmt"
)

func reverseToArray(n int64) []int {
	var result []int

	for n > 0 {
		result = append(result, int(n%10))
		n /= 10
	}
	return result
}

func main() {
	n := int64(12345)
	fmt.Println(reverseToArray(n)) // Output: [5, 4, 3, 2, 1]
}
