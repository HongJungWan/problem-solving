// TODO: 문자열 압축 문제
// TODO: 연속된 동일 문자를, 문자와 그 개수로 대체 -> 압축

package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	newChars := compress(chars)
	fmt.Printf("Compressed array: %s\n", newChars)
}

func compress(chars []byte) []byte {
	n := len(chars)
	if n == 0 {
		return chars
	}

	newChars := make([]byte, 0, n)
	count := 1

	for i := 1; i < n; i++ {
		// 이전 문자와 현재 문자가 동일하면, count 증가
		if chars[i] == chars[i-1] {
			count++
		} else {
			// 이전 문자와 현재 문자가 다르면, 문자와 그 개수를 추가
			newChars = append(newChars, chars[i-1])
			if count > 1 {
				newChars = append(newChars, []byte(strconv.Itoa(count))...)
			}

			// 추가 후 count 초기화
			count = 1
		}
	}

	// 마지막 문자와 그 개수를 추가
	newChars = append(newChars, chars[n-1])
	if count > 1 {
		newChars = append(newChars, []byte(strconv.Itoa(count))...)
	}

	return newChars
}

/*
Go 언어에서 문자열을 인덱스를 통해 접근할 때는 각 문자에 해당하는 바이트 값이 반환.
이는 문자열이 내부적으로 바이트 슬라이스로 표현되기 때문이다.
*/
