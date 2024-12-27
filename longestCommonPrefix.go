package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := ""
	pos := 0
	loop := true
	for loop {
		var letter string
		for i := 0; i < len(strs); i++ {
			if letter == "" {
				letter = string(strs[i][pos])
			}

			if pos >= len(strs[i]) || string(strs[i][pos]) != letter {
				loop = false
				break
			}
		}
		if loop {
			pos++
			prefix = strings.Join([]string{prefix, letter}, "")
		}
	}

	return prefix
}
