// TODO: 로프
// https://www.acmicpc.net/problem/2217

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	ropes := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		ropes[i], _ = strconv.Atoi(scanner.Text())
	}

	// 내림차순 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(ropes)))

	maxWeight := 0
	for i, weight := range ropes {
		// k개의 로프를 사용했을 때 최대 중량 계산 (모든 로프가 해당 중량을 나눠 들어야 함)
		load := weight * (i + 1)
		if load > maxWeight {
			maxWeight = load
		}
	}

	fmt.Println(maxWeight)
}
