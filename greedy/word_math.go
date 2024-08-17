// TODO: 단어 수학
// https://www.acmicpc.net/problem/1339

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	weight := make(map[rune]int)

	// 각 단어를 스캔하면서 각 글자의 자릿수 가치를 계산
	for i := 0; i < N; i++ {
		scanner.Scan()
		word := scanner.Text()
		l := len(word)
		for j, ch := range word {
			// 10의 거듭제곱 자릿수 가치 계산
			weight[ch] += pow(10, l-j-1)
		}
	}

	// 자릿수 가치를 기준으로 정렬하기 위해 슬라이스로 변환
	weights := make([]int, 0, len(weight))
	for _, w := range weight {
		weights = append(weights, w)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	// 가장 큰 자릿수 가치부터 9, 8, ..., 1, 0을 할당
	result := 0
	num := 9
	for _, w := range weights {
		result += w * num
		num--
	}

	fmt.Println(result)
}

// 거듭제곱 계산 함수
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
