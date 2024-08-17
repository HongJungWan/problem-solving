// 대표값 2
// https://www.acmicpc.net/problem/2587

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Numbers []int

func (numbers Numbers) Len() int {
	return len(numbers)
}

func (numbers Numbers) Less(i, j int) bool {
	// 사전순으로 정렬
	return numbers[i] < numbers[j]
}

func (numbers Numbers) Swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

func main() {
	avg := 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := 5
	numbers := make(Numbers, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &numbers[i])
	}

	// 정렬
	sort.Sort(numbers)

	for i := 0; i < len(numbers); i++ {
		avg += numbers[i]
	}
	avg /= n

	fmt.Println(avg)
	fmt.Println(numbers[len(numbers)/2])
}
