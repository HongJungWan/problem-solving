// TODO: 붙임성 좋은 총총이
// https://www.acmicpc.net/problem/26069

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scan(&N)
	scanner := bufio.NewScanner(os.Stdin)

	// 무지개 댄스를 춘 사람들을 추적하는 데 사용되는 맵
	dancing := make(map[string]bool)
	dancing["ChongChong"] = true // 총총이는 시작부터 무지개 댄스를 춥니다.

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		var A, B string
		fmt.Sscanf(line, "%s %s", &A, &B)

		// A가 무지개 댄스를 추고 있다면 B도 무지개 댄스를 추게 됩니다.
		if dancing[A] {
			dancing[B] = true
		}

		// 반대로, B가 무지개 댄스를 추고 있다면 A도 무지개 댄스를 추게 됩니다.
		if dancing[B] {
			dancing[A] = true
		}
	}
	fmt.Println(len(dancing))
}
