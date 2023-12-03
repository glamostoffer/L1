package main

import (
	"fmt"
	"strings"
	"unicode"
)

func easierReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWords(s string) string {
	str := []rune(s)
	var temp []rune
	var words [][]rune
	for i := 0; i < len(str); i++ {
		if unicode.IsSpace(str[i]) {
			words = append(words, temp)
			temp = []rune{}
		} else {
			temp = append(temp, str[i])
		}
	}

	words = append(words, temp)
	s = ""

	for i := len(words) - 1; i >= 0; i-- {
		s += string(words[i]) + " "
	}

	s = s[:len(s)-1]

	return s
}

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(reverseWords(s))
	fmt.Println(easierReverseWords(s))
}
