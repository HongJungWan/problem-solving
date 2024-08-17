// TODO: 임진왜란
// https://www.acmicpc.net/problem/3077

package main

import (
	"fmt"
)

func readInput(N int, correctOrder map[string]int, userOrder []string) {
	for i := 0; i < N; i++ {
		var battle string
		fmt.Scan(&battle)
		correctOrder[battle] = i
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&userOrder[i])
	}
}

func calculateScore(N int, correctOrder map[string]int, userOrder []string) (int, int) {
	score := 0
	totalPairs := N * (N - 1) / 2
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if correctOrder[userOrder[i]] < correctOrder[userOrder[j]] {
				score++
			}
		}
	}
	return score, totalPairs
}

func main() {
	var N int
	fmt.Scan(&N)

	correctOrder := make(map[string]int)
	userOrder := make([]string, N)

	readInput(N, correctOrder, userOrder)
	score, totalPairs := calculateScore(N, correctOrder, userOrder)

	fmt.Printf("%d/%d\n", score, totalPairs)
}
