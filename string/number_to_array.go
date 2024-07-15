// TODO: 자연수 n을 각 자리 숫자를 원소로 가지는 배열 형태로 리턴

package main

import (
	"fmt"
	"strconv"
)

func toArray(n int64) []int {
	str := strconv.FormatInt(n, 10)
	result := make([]int, len(str))

	for i, char := range str {
		result[i] = int(char - '0')
	}
	return result
}

func main() {
	n := int64(12345)
	fmt.Println(toArray(n)) // Output: [1, 2, 3, 4, 5]
}

/*
strconv.FormatInt 함수를 사용하여 정수 n을 문자열로 변환,

int(char - '0')는 문자를 해당 숫자 값으로 변환.
*/
