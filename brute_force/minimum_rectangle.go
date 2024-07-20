// TODO: 최소 직사각형
// TODO: 가장 큰 길이를 가로로 변경
// https://school.programmers.co.kr/learn/courses/30/lessons/86491

package main

import "fmt"

func minWalletSize(sizes [][]int) int {
	maxWidth := 0
	maxHeight := 0

	for _, size := range sizes {
		w, h := size[0], size[1]
		if w < h {
			w, h = h, w
		}
		if w > maxWidth {
			maxWidth = w
		}
		if h > maxHeight {
			maxHeight = h
		}
	}

	return maxWidth * maxHeight
}

func main() {
	sizes1 := [][]int{{60, 50}, {30, 70}, {60, 30}, {80, 40}}
	sizes2 := [][]int{{10, 7}, {12, 3}, {8, 15}, {14, 7}, {5, 15}}
	sizes3 := [][]int{{14, 4}, {19, 6}, {6, 16}, {18, 7}, {7, 11}}

	fmt.Println(minWalletSize(sizes1)) // 4000
	fmt.Println(minWalletSize(sizes2)) // 120
	fmt.Println(minWalletSize(sizes3)) // 133
}

/*
"하지만 2번 명함을 가로로 눕혀 수납한다면 80(가로) x 50(세로) 크기의 지갑으로 모든 명함들을 수납할 수 있습니다."

이 문장에서 명함을 회전시켜 더 큰 길이를 가로로 맞추는 것이 필요함을 알 수 있다.
*/
