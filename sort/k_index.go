// TODO: K번째 수
// https://school.programmers.co.kr/learn/courses/30/lessons/42748

package main

import (
	"fmt"
	"sort"
)

func kIndex(array []int, commands [][]int) []int {
	var answer []int

	for _, command := range commands {
		i, j, k := command[0], command[1], command[2]

		slice := make([]int, j-i+1)
		copy(slice, array[i-1:j])

		sort.Ints(slice)

		answer = append(answer, slice[k-1])
	}

	return answer
}

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}

	fmt.Println(kIndex(array, commands)) // [5, 6, 3]
}
