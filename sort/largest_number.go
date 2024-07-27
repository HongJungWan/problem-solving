// TODO: 가장 큰 수
// https://school.programmers.co.kr/learn/courses/30/lessons/42746

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ByLargestNumber []string

func (byLargestNumber ByLargestNumber) Len() int { return len(byLargestNumber) }

func (byLargestNumber ByLargestNumber) Swap(i, j int) {
	byLargestNumber[i], byLargestNumber[j] = byLargestNumber[j], byLargestNumber[i]
}

func (byLargestNumber ByLargestNumber) Less(i, j int) bool {
	return byLargestNumber[i]+byLargestNumber[j] > byLargestNumber[j]+byLargestNumber[i]
}

func largestNumber(numbers []int) string {
	strNumbers := make([]string, len(numbers))
	for i, num := range numbers {
		strNumbers[i] = strconv.Itoa(num)
	}

	sort.Sort(ByLargestNumber(strNumbers))

	result := strings.Join(strNumbers, "")

	if result[0] == '0' {
		return "0"
	}

	return result
}

func main() {
	numbers1 := []int{6, 10, 2}
	fmt.Println(largestNumber(numbers1)) // Output: "6210"

	numbers2 := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(numbers2)) // Output: "9534330"
}
