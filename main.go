package main

import (
	"fmt"
	"sort"
)

/***
	Main method to identify if cards are straight
***/
func IsStraight(cards []int) bool {
	done := initialValidate(cards)
	if !done {
		return false
	}

	var straight []int
	cards = sortCards(cards)

	i := 0
	for i < len(cards) {
		if i == (len(cards) - 1) {
			break
		}

		current := cards[i]
		next := cards[i+1]

		if len(straight) == 5 {
			break
		}

		if validateStraight(current, next) {
			straight = append(straight, current)

			if len(straight) == 5 {
				break
			}

			if i == (len(cards) - 2) {
				straight = append(straight, next)
			}

			i++
			continue
		}

		straight = []int{}
		i++
	}

	if len(straight) == 5 {
		return true
	}

	return false
}

/***
	Validates the initial conditions to be a straight
***/
func initialValidate(cards []int) bool {
	if len(cards) < 4 {
		return false
	}

	if len(cards) > 7 {
		return false
	}

	return true
}

/***
	Sort the cards ascending
***/
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

/***
	Valid that the next card is the one indicated
***/
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


func main()  {
	if IsStraight([]int{2,7,8,5,10,9,11}) {
		fmt.Println("IsStraight")
		return
	}

	fmt.Println("Not IsStraight")
}