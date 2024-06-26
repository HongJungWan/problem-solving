// TODO: 주식 매매 최적 시기

/*
문제 설명:
주어진 주식 가격 배열에서 주식을 사고 팔 수 있는 최대의 이익을 계산하는 문제.
하루에 주식을 여러 번 거래할 수 있지만, 동시에 두 개 이상의 주식을 보유할 수는 없다.
연속된 날의 주식 가격이 증가하면, 그 구간을 모두 사서 파는 것이 이익을 극대화하는 방법.

예시 입력: prices = [7, 1, 5, 3, 6, 4]

예시 출력: 7

[아이디어]
주식 가격이 오르는 구간에서만 이익을 취하는 것.
주식을 사고 파는 모든 가능한 경우를 탐색하지 않고도 최대 이익을 얻을 수 있는 Greedy 접근법
*/

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum profit:", maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
