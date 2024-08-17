// TODO: 같은 숫자는 싫어
// https://school.programmers.co.kr/learn/courses/30/lessons/12906

package main

import "fmt"

func removeDuplicates(arr []int) []int {
	var stack []int

	for _, num := range arr {
		if len(stack) == 0 || stack[len(stack)-1] != num {
			stack = append(stack, num)
		}
	}
	return stack
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 3, 3, 0, 1, 1})) // [1, 3, 0, 1]
	fmt.Println(removeDuplicates([]int{4, 4, 4, 3, 3}))       // [4, 3]
}
