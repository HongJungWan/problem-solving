// TODO: 완주하지 못한 선수
// https://school.programmers.co.kr/learn/courses/30/lessons/42576

package main

import "fmt"

func findIncompleteParticipant(participant []string, completion []string) string {
	participantCount := make(map[string]int)

	for _, p := range participant {
		participantCount[p]++
	}

	for _, c := range completion {
		participantCount[c]--
	}

	// key, value 반환
	for participantItem, count := range participantCount {
		if count > 0 {
			return participantItem
		}
	}

	return ""
}

func main() {
	participants1 := []string{"leo", "kiki", "eden"}
	completion1 := []string{"eden", "kiki"}
	fmt.Println(findIncompleteParticipant(participants1, completion1)) // "leo"

	participants2 := []string{"marina", "josipa", "nikola", "vinko", "filipa"}
	completion2 := []string{"josipa", "filipa", "marina", "nikola"}
	fmt.Println(findIncompleteParticipant(participants2, completion2)) // "vinko"

	participants3 := []string{"mislav", "stanko", "mislav", "ana"}
	completion3 := []string{"stanko", "ana", "mislav"}
	fmt.Println(findIncompleteParticipant(participants3, completion3)) // "mislav"
}
