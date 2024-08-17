// TODO: 좌표 정렬하기 2
// https://www.acmicpc.net/problem/11651

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Point 구조체 정의: 각 점의 x, y 좌표를 저장
type Point struct {
	x, y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N) // 점의 개수 N 입력 받기

	points := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &points[i].x, &points[i].y) // 각 점의 좌표 입력 받기
	}

	// points 슬라이스를 y 좌표가 증가하는 순으로 정렬하되,
	// ‼ y 좌표가 같을 경우 x 좌표가 증가하는 순서로 정렬
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	// 정렬된 점들 출력
	for _, p := range points {
		fmt.Fprintln(writer, p.x, p.y)
	}
}

// sort.Slice 함수: 슬라이스를 정렬하는 데 사용.
// sort.Slice 함수: 정렬하려는 슬라이스와 두 요소를 비교하는 함수를 인자로 받는다.
//
// 비교 함수는 두 인덱스 i와 j를 매개변수로 하며,
// 슬라이스의 i번째 요소가 j번째 요소보다 "먼저 와야 한다"
// (즉, 더 작다)는 의미에서 true를 반환, 그렇지 않다면 false
