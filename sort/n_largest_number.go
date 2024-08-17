// N번째 큰 수
// https://www.acmicpc.net/problem/2693

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var N int
	fmt.Fscanln(os.Stdin, &N)

	for i := 0; i < N; i++ {
		arr := make([]int, 10) // 각 테스트 케이스마다 새로운 배열 생성

		for j := 0; j < 10; j++ {
			fmt.Scan(&arr[j])
		}
		sort.Ints(arr)
		fmt.Println(arr[7])
	}
}
