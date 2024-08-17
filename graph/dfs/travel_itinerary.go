// 여행 경로
// https://school.programmers.co.kr/learn/courses/30/lessons/43164

package main

import (
	"fmt"
	"sort"
)

func FindItineraryDFS(graph map[string][]string, start string, path []string, ticketNum int) []string {
	if len(path) == ticketNum+1 {
		return path
	}

	for i, city := range graph[start] {
		// Remove city from graph
		graph[start] = append(graph[start][:i], graph[start][i+1:]...)
		newPath := append(path, city)
		result := FindItineraryDFS(graph, city, newPath, ticketNum)

		// Add city back to graph
		graph[start] = append(graph[start][:i], append([]string{city}, graph[start][i:]...)...)

		if result != nil {
			return result
		}
	}
	return nil
}

func FindItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	for start := range graph {
		sort.Strings(graph[start]) // 오름차순 정렬
	}
	return FindItineraryDFS(graph, "ICN", []string{"ICN"}, len(tickets))
}

func main() {
	tickets1 := [][]string{{"ICN", "JFK"}, {"HND", "IAD"}, {"JFK", "HND"}}
	tickets2 := [][]string{{"ICN", "SFO"}, {"ICN", "ATL"}, {"SFO", "ATL"}, {"ATL", "ICN"}, {"ATL", "SFO"}}
	fmt.Println("Itinerary 1:", FindItinerary(tickets1))
	fmt.Println("Itinerary 2:", FindItinerary(tickets2))
}
