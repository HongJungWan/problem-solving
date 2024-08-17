// TODO: 거스름돈
// https://www.acmicpc.net/problem/2720

package main

import (
	"fmt"
)

func main() {
	var T, C int
	fmt.Scan(&T)

	denominations := []int{25, 10, 5, 1}

	for i := 0; i < T; i++ {
		fmt.Scan(&C)
		result := make([]int, 4)

		for j, coin := range denominations {
			result[j] = C / coin
			C %= coin
		}

		fmt.Printf("%d %d %d %d\n", result[0], result[1], result[2], result[3])
	}
}
