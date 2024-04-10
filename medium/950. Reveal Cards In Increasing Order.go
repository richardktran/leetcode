package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	if n == 1 {
		return deck
	}
	sort.Ints(deck)
	result := make([]int, n)
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i
	}

	for _, card := range deck {
		result[queue[0]] = card
		queue = queue[1:]
		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}

	return result
}

func deckRevealedIncreasingTest() {
	println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7})) // [2, 13, 3, 11, 5, 17, 7]
	println(deckRevealedIncreasing([]int{1, 1000}))                // [1,1000]
}

// func main() {
// 	deckRevealedIncreasingTest()
// }
