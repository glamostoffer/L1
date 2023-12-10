package main

import (
	"fmt"
	"unicode"
)

func isUniqueChars(s string) bool {
	str := []rune(s)
	symbols := make(map[rune]bool)
	for _, val := range str {
		if symbols[unicode.ToLower(val)] {
			return false
		}
		symbols[unicode.ToLower(val)] = true
	}
	return true
}

// разница между е-Е 32

func main() {
	fmt.Println(isUniqueChars("true? yas"))
	fmt.Println(isUniqueChars("false? yeS"))
	fmt.Println(isUniqueChars("  "))
}
