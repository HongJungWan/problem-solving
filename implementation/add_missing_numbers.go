// TODO: 없는 숫자 더하기
// https://school.programmers.co.kr/learn/courses/30/lessons/86051

package main

import "fmt"

func addMissingNumber(numbers []int) int {
	compare := make(map[int]bool, 10)
	var result int

	for _, number := range numbers {
		compare[number] = true
	}

	for i := 0; i <= 9; i++ {
		if !compare[i] {
			result += i
		}
	}

	return result
}

func main() {
	numbersOne := []int{1, 2, 3, 4, 6, 7, 8, 0}
	numbersTwo := []int{5, 8, 4, 0, 6, 7, 9}

	fmt.Println(addMissingNumber(numbersOne)) // 14
	fmt.Println(addMissingNumber(numbersTwo)) // 6
}
