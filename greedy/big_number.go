// TODO: 큰 수 만들기
// TODO: https://school.programmers.co.kr/learn/courses/30/lessons/42883

package main

import (
	"fmt"
)

func bigNumberMake(number string, k int) string {
	stack := []rune{}
	toRemove := k

	for _, digit := range number {
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-toRemove]

	return string(stack)
}

func main() {
	// 예제 테스트
	fmt.Println(bigNumberMake("1924", 2))       // "94"
	fmt.Println(bigNumberMake("1231234", 3))    // "3234"
	fmt.Println(bigNumberMake("4177252841", 4)) // "775841"
}
