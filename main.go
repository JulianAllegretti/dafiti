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

		if !validateStraight(current, next) {
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

func validateStraight(currentCard int, nextCard int) bool{
	if currentCard > 14 || nextCard > 14 {
		return false
	}

	if currentCard < 2 || nextCard < 2 {
		return false
	}

	if currentCard == 14 {
		currentCard = 1
	}

	if nextCard == 14 {
		nextCard = 1
	}

	if (currentCard + 1) != nextCard {
		return false
	}

	return true
}
