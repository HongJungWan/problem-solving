// TODO: 별 찍기 5
// https://www.acmicpc.net/problem/2442

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // N 입력 받기
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	for i := 1; i <= N; i++ {
		// 각 줄마다 필요한 공백 출력
		for space := 0; space < N-i; space++ {
			fmt.Print(" ")
		}

		// 별 출력 (2*i-1 개)
		for star := 1; star <= 2*i-1; star++ {
			fmt.Print("*")
		}

		// 한 줄이 끝나면 개행
		fmt.Println()
	}
}
