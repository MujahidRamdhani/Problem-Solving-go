package main

import (
	"fmt"
	"strings"
)

// fungsi untuk memeriksa apakah string memiliki tanda kurung yang seimbang
func IsBalanced(s string) string {
	var stack []rune
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if char == ' ' {
			continue 
		}

		switch char {
			
		case '(', '{', '[':
			stack = append(stack, char)

		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return "NO"
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	tests := []string{
		"{ [ ( ) ] }",
		"{ [ ( ] ) }",
		"{ ( ( [ ] ) [ ] ) [ ] }",
	}

	for _, test := range tests {
		cleanedInput := strings.ReplaceAll(test, " ", "")
		fmt.Printf("Input: %s\nOutput: %s\n\n", test, IsBalanced(cleanedInput))
	}
}
