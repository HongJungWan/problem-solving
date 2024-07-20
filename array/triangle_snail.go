// TODO: 삼각 달팽이
// TODO: 주어진 삼각형 배열을 달팽이 패턴으로 채워서 1차원 배열로 반환.
// https://school.programmers.co.kr/learn/courses/30/lessons/68645?language=java

package main

import "fmt"

func generateSnailTriangle(n int) []int {
	// 삼각형 달팽이 배열 생성
	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	num := 1
	// x와 y를 초기화, x = -1, y = 0로 시작하는 이유는 첫 이동이 아래로 한 칸 내려가면서 시작되기 때문이다.
	x, y := -1, 0

	// 방향 벡터 (아래, 오른쪽, 위쪽)
	dx := []int{1, 0, -1}
	dy := []int{0, 1, -1}
	dir := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			x += dx[dir]
			y += dy[dir]
			triangle[x][y] = num

			num++
		}

		// 이동 후, 방향을 바꾼다.
		// 방향은 세 가지 (아래, 오른쪽, 위쪽)이며, 순서대로 순환해야 한다.
		// dir을 0, 1, 2로 사용하여 순환적으로 이동한다.
		// 이를 위해, dir을 1씩 증가시키고 3으로 나눈 나머지를 구한다. 이렇게 하면 dir이 0, 1, 2로 순환한다.
		dir = (dir + 1) % 3
	}

	// 결과 배열로 변환
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			result = append(result, triangle[i][j])
		}
	}

	return result
}

func main() {
	fmt.Println(generateSnailTriangle(4)) // [1,2,9,3,10,8,4,5,6,7]
	fmt.Println(generateSnailTriangle(5)) // [1,2,12,3,13,11,4,14,15,10,5,6,7,8,9]
	fmt.Println(generateSnailTriangle(6)) // [1,2,15,3,16,14,4,17,21,13,5,18,19,20,12,6,7,8,9,10,11]
}
