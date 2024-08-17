// TODO: 패션왕 신해빈
// https://www.acmicpc.net/problem/9375

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 의상 Input.
func inputClothes(t int, reader *bufio.Reader, writer *bufio.Writer) {
	for ; t > 0; t-- {
		var n int
		fmt.Fscanln(reader, &n)

		clothes := make(map[string]int)
		for i := 0; i < n; i++ {
			var name, category string
			fmt.Fscanln(reader, &name, &category)
			clothes[category]++
		}
		result := calculateCombinations(clothes)
		fmt.Fprintln(writer, result)
	}
}

// 주어진 의상들의 조합 수를 계산.
func calculateCombinations(clothes map[string]int) int {
	result := 1
	for _, count := range clothes {
		result *= (count + 1) // 각 카테고리별 의상 수 + 1 (안 입는 경우)
	}
	return result - 1 // 모두 안 입는 경우 1 빼기
}

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	inputClothes(t, reader, writer)
}
