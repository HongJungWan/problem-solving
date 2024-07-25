// TODO: 두 개 뽑아서 더하기
// https://school.programmers.co.kr/learn/courses/30/lessons/68644

package main

import (
	"fmt"
	"sort"
)

func pickTwoAdd(numbers []int) []int {
	resultSet := make(map[int]bool)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]
			resultSet[sum] = true
		}
	}

	// 결과를 슬라이스로 변환하고 정렬
	result := make([]int, 0, len(resultSet))
	for key := range resultSet {
		result = append(result, key)
	}
	sort.Ints(result)

	return result
}

func main() {
	numbers1 := []int{2, 1, 3, 4, 1}
	numbers2 := []int{5, 0, 2, 7}

	fmt.Println(pickTwoAdd(numbers1)) // [2 3 4 5 6 7]
	fmt.Println(pickTwoAdd(numbers2)) // [2 5 7 9 12]
}
