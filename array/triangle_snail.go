// TODO: 삼각 달팽이 [이해 못함]
// TODO: 주어진 삼각형 배열을 달팽이 패턴으로 채워서 1차원 배열로 반환.
// https://school.programmers.co.kr/learn/courses/30/lessons/68645?language=java

package main

import "fmt"

// 삼각형 배열을 달팽이 패턴으로 채워서 1차원 배열로 반환하는 함수
func generateSnailTriangle(n int) []int {
	triangle := createTriangle(n)
	fillSnailPattern(triangle, n)
	return flattenTriangle(triangle)
}

// 삼각형 배열 생성
func createTriangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}
	return triangle
}

// 삼각형 배열을 달팽이 패턴으로 채우는 함수
func fillSnailPattern(triangle [][]int, n int) {
	num := 1

	x, y := -1, 0
	directions := [][]int{{1, 0}, {0, 1}, {-1, -1}}
	dir := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			x += directions[dir][0]
			y += directions[dir][1]

			triangle[x][y] = num
			num++
		}
		dir = (dir + 1) % 3
	}
}

// 2차원 삼각형 배열을 1차원 배열로 변환하는 함수
func flattenTriangle(triangle [][]int) []int {
	result := make([]int, 0)

	// 2차원 배열의 각 행을 반복
	for i := range triangle {
		result = append(result, triangle[i]...)
	}
	return result
}

func main() {
	fmt.Println(generateSnailTriangle(4)) // Output: [1 2 9 3 10 8 4 5 6 7]
	fmt.Println(generateSnailTriangle(5)) // Output: [1 2 12 3 13 11 4 14 15 10 5 6 7 8 9]
	fmt.Println(generateSnailTriangle(6)) // Output: [1 2 15 3 16 14 4 17 21 13 5 18 19 20 12 6 7 8 9 10 11]
}

/*
i = 0: triangle[0] = make([]int, 1) -> [[0], nil, nil, nil]
i = 1: triangle[1] = make([]int, 2) -> [[0], [0, 0], nil, nil]
i = 2: triangle[2] = make([]int, 3) -> [[0], [0, 0], [0, 0, 0], nil]
i = 3: triangle[3] = make([]int, 4) -> [[0], [0, 0], [0, 0, 0], [0, 0, 0, 0]]
*/
