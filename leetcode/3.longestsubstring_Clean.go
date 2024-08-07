package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var lengthStr int = len(s)
	max_int := 1
	if lengthStr == 0 {
		max_int = 0
	}
	currIdx := 0
	var arrayStreak []string
	currLongestStreak := ""
	for true {
		// fmt.Printf("\n\n########Current loop-%v\n", currIdx)
		for i := currIdx; i < lengthStr; i++ {
			tmpStr := s[i : i+1]
			tmpInt := strings.Count(currLongestStreak, tmpStr)
			if tmpInt > 0 {
				arrayStreak = append(arrayStreak, currLongestStreak)
				currLongestStreak = ""
				break
			} else {
				currLongestStreak += tmpStr
			}
			// fmt.Printf("Current string:%v\n", tmpStr)
		}
		currIdx++

		if currIdx >= lengthStr {
			break
		}
	}
	fmt.Println(arrayStreak)

	for idx := range arrayStreak {
		if max_int < len(arrayStreak[idx]) {
			max_int = len(arrayStreak[idx])
		}
	}
	return max_int
}
func main() {
	s1 := " "
	jum := lengthOfLongestSubstring(s1)

	fmt.Printf("Jum:%v", jum)
}
