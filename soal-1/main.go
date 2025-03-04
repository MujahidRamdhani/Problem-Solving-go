package main

import (
	"fmt"
	"strings"
)

// fungsi untuk menghasilkan deret A000124
func GenerateA000124Sequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = (i * (i + 1) / 2) + 1
	}
	return sequence
}

// fungsi mengubah slice integer menjadi string 
func FormatSequence(sequence []int) string {
	strValues := make([]string, len(sequence))
	for i, num := range sequence {
		strValues[i] = fmt.Sprintf("%d", num)
	}
	return strings.Join(strValues, "-")
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Input harus bilangan positif!")
		return
	}

	sequence := GenerateA000124Sequence(n)
	formattedOutput := FormatSequence(sequence)
	fmt.Println("Output:", formattedOutput)
}
