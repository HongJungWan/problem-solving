// TODO: 배열 돌리기
// TODO: 이차원 배열의 인덱스 변환 규칙을 알아야 한다.
// TODO: 손이던 머리로던 시뮬레이션 하고, 공식 도출해야 함 -> 회전 시 각 원소의 새로운 위치를 계산하는 방법을 이해

package main

import "fmt"

const (
	N = 4
	M = 4
)

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println("Original Array:")
	printArray(arr)

	// 배열 회전
	rotatedArr := rotateArray(arr)

	fmt.Println("\n\nRotated Array:")
	printArray(rotatedArr)
}

func rotateArray(arr [][]int) [][]int {
	newArr := make([][]int, N)
	for i := range newArr {
		newArr[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			newArr[N-j-1][i] = arr[i][j]
		}
	}
	return newArr
}

func printArray(arr [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

// [Original Array]
//    1 2 3 4
//    5 6 7 8
//    9 10 11 12
//    13 14 15 16

// [Rotated Array]
//    4 8 12 16
//    3 7 11 15
//    2 6 10 14
//    1 5 9 13

/* 아이디어

원래 배열의 각 좌표와 회전 후 좌표를 비교해 보면,

(0, 0) → (3, 0)
(0, 1) → (2, 0)
(0, 2) → (1, 0)
(0, 3) → (0, 0)
(1, 0) → (3, 1)
(1, 1) → (2, 1)
(1, 2) → (1, 1)
(1, 3) → (0, 1)

(i, j) → (N - j - 1, i)라는 규칙을 도출
*/
