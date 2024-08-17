// TODO: 카드 1
// https://www.acmicpc.net/problem/2161

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// simulateCardGame은 N장의 카드로 카드 게임을 시뮬레이션하고, 버려진 카드 목록과 마지막 카드를 반환
func simulateCardGame(N int) ([]int, int) {
	queue := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		queue = append(queue, i)
	}

	discarded := make([]int, 0)

	// while 반복문 구문 수정
	for len(queue) > 1 {
		// 가장 위의 카드를 버림
		discarded = append(discarded, queue[0])
		queue = queue[1:]

		// 새로운 가장 위의 카드를 아래로 옮김
		queue = append(queue, queue[0])
		queue = queue[1:]
	}

	// 마지막 카드를 반환
	lastCard := queue[0]
	return discarded, lastCard
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 읽기
	input, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(input[:len(input)-1])

	// 게임 시뮬레이션 실행
	discarded, lastCard := simulateCardGame(N)

	// 버려진 카드 출력
	for _, card := range discarded {
		fmt.Fprintf(writer, "%d ", card)
	}

	// 마지막으로 남은 카드 출력
	fmt.Fprintf(writer, "%d\n", lastCard)
}
