// TODO: 별 찍기 9
// https://www.acmicpc.net/problem/2446

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printStar(n, max int) {
	space := (max - n) / 2
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // N 입력 받기
	N, _ := strconv.Atoi(scanner.Text())

	// 별 감소 부분
	for i := 2*N - 1; i > 0; i -= 2 {
		printStar(i, 2*N-1)
	}

	// 별 증가 부분 (첫 줄 제외)
	for i := 3; i <= 2*N-1; i += 2 {
		printStar(i, 2*N-1)
	}
}
