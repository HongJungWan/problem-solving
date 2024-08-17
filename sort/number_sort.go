// TODO: 수 정렬하기
// https://www.acmicpc.net/problem/2750

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N) // N 개수 입력 받기

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &numbers[i])
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	for _, number := range numbers {
		fmt.Fprintln(writer, number)
	}
}
