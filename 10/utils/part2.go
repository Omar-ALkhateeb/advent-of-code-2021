package utils

import (
	"fmt"
	"sort"
	"strings"
)

func AutoComplete(subMap string) int {
	// fmt.Println(subMap)
	lines := strings.Split(subMap, "\n")
	scores := []int{}

	m := make(map[string]string)
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"
	m["<"] = ">"
	p := make(map[string]int)
	p["["] = 2
	p["{"] = 3
	p["("] = 1
	p["<"] = 4
	// remove corrupted lines
loop:
	for pos := 0; pos < len(lines); pos++ {
		toClose := ""
		flag := false
		for _, char := range lines[pos] {
			if char == '[' || char == '(' || char == '{' || char == '<' {
				toClose += string(char)
			}
			if char == ']' || char == ')' || char == '}' || char == '>' {
				if string(char) == m[string(toClose[len(toClose)-1])] {
					toClose = toClose[:len(toClose)-1]
				} else {
					flag = true
					break
				}
			}
		}
		if flag {
			lines = append(lines[:pos], lines[pos+1:]...)
			pos--
			continue loop
		}
	}

	for pos, line := range lines {
		toClose := ""
		for _, char := range line {
			if char == '[' || char == '(' || char == '{' || char == '<' {
				toClose += string(char)
			}
			if char == ']' || char == ')' || char == '}' || char == '>' {
				if string(char) == m[string(toClose[len(toClose)-1])] {
					toClose = toClose[:len(toClose)-1]
				}
			}
		}
		// fmt.Println(toClose)
		scores = append(scores, 0)
		for i := len(toClose) - 1; i >= 0; i-- {
			char := toClose[i]
			scores[pos] *= 5
			scores[pos] += p[string(char)]
		}
	}

	middleIndex := len(scores) / 2

	sort.Ints(scores)

	fmt.Println(scores)

	return scores[middleIndex]

}
