// TODO: 알파벳 찾기
// https://www.acmicpc.net/problem/10809

package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	positions := make([]int, 26)
	for i := range positions {
		positions[i] = -1
	}

	for i, char := range S {
		index := char - 'a'
		if positions[index] == -1 {
			positions[index] = i
		}
	}

	for _, pos := range positions {
		fmt.Printf("%d ", pos)
	}
}
