// TODO: H-Index
// https://school.programmers.co.kr/learn/courses/30/lessons/42747

package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)

	answer := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > answer {
			answer++
		} else {
			break
		}
	}

	return answer
}

func main() {
	citations := []int{3, 0, 6, 1, 5}

	fmt.Println(hIndex(citations))
}
