package main

import (
	"fmt"
	"sort"
)

func main()  {
	if IsStraight([]int{14, 3, 2, 6, 4 ,5}) {
		fmt.Println("IsStraight")
	}
}

func IsStraight(cards []int) bool {
	if len(cards) < 4 {
		return false
	}

	if len(cards) > 7 {
		return false
	}

	cards = sortCards(cards)

	i := 0
	for i < len(cards) {
		if i == (len(cards) - 1) {
			break
		}

		current := cards[i]
		next := cards[i+1]

		if current > 14 || next > 14 {
			return false
		}

		if current < 2 || next < 2 {
			return false
		}

		if current == 14 {
			current = 1
		}

		if next == 14 {
			next = 1
		}

		if current+1 != next {
			return false
		}

		i++
	}

	return true
}

func sortCards(cards []int) []int{
	sort.Slice(cards, func(i, j int) bool {
		if cards[j] == 14 {
			return false
		}

		if cards[i] == 14 {
			return true
		}

		return cards[i] < cards[j]
	})

	return cards
}
