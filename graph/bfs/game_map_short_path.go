// 게임 맵 최단거리
// https://school.programmers.co.kr/learn/courses/30/lessons/1844

package main

import (
	"fmt"
)

type Position struct {
	x, y, distance int
}

func isValid(x, y, n, m int, maps [][]int, visited [][]bool) bool {
	return x >= 0 && x < n && y >= 0 && y < m && maps[x][y] == 1 && !visited[x][y]
}

func shortPathDFS(maps [][]int) int {
	n := len(maps)
	m := len(maps[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	// Directions: North, East, South, West
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	queue := []Position{{0, 0, 1}}
	visited[0][0] = true

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos.x == n-1 && pos.y == m-1 {
			return pos.distance
		}

		for i := 0; i < 4; i++ {
			nextX := pos.x + dx[i]
			nextY := pos.y + dy[i]

			if isValid(nextX, nextY, n, m, maps, visited) {
				visited[nextX][nextY] = true
				queue = append(queue, Position{nextX, nextY, pos.distance + 1})
			}
		}
	}
	return -1
}

func main() {
	mapsSolutionOne := [][]int{
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{0, 0, 0, 0, 1},
	}
	fmt.Println(shortPathDFS(mapsSolutionOne)) // 11

	mapSolutionTwo := [][]int{
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1},
	}
	fmt.Println(shortPathDFS(mapSolutionTwo)) // -1
}
