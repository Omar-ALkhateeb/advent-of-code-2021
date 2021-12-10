package utils

import (
	"strings"
)

func FindErrors(subMap string) int {
	// fmt.Println(subMap)
	lines := strings.Split(subMap, "\n")
	toClose := ""
	// lineError := []string{}
	m := make(map[string]string)
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"
	m["<"] = ">"
	p := make(map[string]int)
	p["]"] = 57
	p["}"] = 1197
	p[")"] = 3
	p[">"] = 25137
	sum := 0
	// fmt.Println(lines)
	for _, line := range lines {
		for _, char := range line {
			if char == '[' || char == '(' || char == '{' || char == '<' {
				toClose += string(char)
			}
			if char == ']' || char == ')' || char == '}' || char == '>' {
				if string(char) == m[string(toClose[len(toClose)-1])] {
					toClose = toClose[:len(toClose)-1]
				} else {
					// fmt.Println(m[string(char)])
					// fmt.Println(p[string(char)])
					// fmt.Println(string(char))
					sum += p[string(char)]

					// lineError = append(lineError, "Expected "+string(toClose[len(toClose)-1])+" but found "+string(char)+" instead.")
					break
				}
			}
			// fmt.Printf("character %c starts at byte position %d\n", char, pos)
		}
	}

	// fmt.Println(lineError)
	// fmt.Println(sum)
	return sum

}
