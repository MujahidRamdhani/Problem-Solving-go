package main

import (
	"fmt"
)

// fungsi mengembalikan daftar peringkat berdasarkan skor yang diberikan
func getDenseRank(scores []int, gitsScores []int) []int {
	uniqueScores := removeDuplicates(scores)
	var ranks []int

	for _, score := range gitsScores {
		rank := getRank(uniqueScores, score)
		ranks = append(ranks, rank)
	}

	return ranks
}

// fungsi menghapus duplikasi dalam daftar skor dan mengembalikannya sebagai slice unik
func removeDuplicates(scores []int) []int {
	uniqueMap := make(map[int]bool)
	var uniqueScores []int

	for _, score := range scores {
		if !uniqueMap[score] {
			uniqueScores = append(uniqueScores, score)
			uniqueMap[score] = true
		}
	}

	return uniqueScores
}

// fungsi menghitung peringkat berdasarkan daftar skor unik
func getRank(uniqueScores []int, score int) int {
	low, high := 0, len(uniqueScores)-1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case uniqueScores[mid] == score:
			return mid + 1
		case uniqueScores[mid] < score:
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return low + 1
}

func main() {
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	gitsScores := []int{5, 25, 50, 120}

	ranks := getDenseRank(scores, gitsScores)
	fmt.Println(ranks) 
}
