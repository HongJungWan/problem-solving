// TODO: 거스름돈
// https://www.acmicpc.net/problem/5585

package main

import (
	"fmt"
)

func main() {
	var price int
	fmt.Scanf("%d", &price)

	change := 1000 - price
	coins := []int{500, 100, 50, 10, 5, 1}
	count := 0

	for _, coin := range coins {
		count += change / coin
		change %= coin
	}

	fmt.Println(count)
}
