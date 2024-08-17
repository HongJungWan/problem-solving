// TODO: 반지
// https://www.acmicpc.net/problem/5555

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 주어진 반지들 중에서 찾고자하는 문자열을 포함하는 반지의 수 계산
func countRingsContainingTarget(target string, rings []string) int {
	count := 0
	for _, ring := range rings {
		ring += ring // 순환 구조 생성
		if strings.Contains(ring, target) {
			count++
		}
	}
	return count
}

func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	// 찾고자 하는 문자열 읽기
	scanner.Scan()
	target := scanner.Text()

	// 반지의 개수 읽기
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// 반지 문자열들 읽기
	rings := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		rings[i] = scanner.Text()
	}

	return target, rings
}

func main() {
	target, rings := readInput()
	count := countRingsContainingTarget(target, rings)
	fmt.Println(count)
}
