package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input file
	// fmt.Println("Args:", os.Args)
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var score int = 0
	var cards string
	var winningCards []string
	var actualCards []string
	for sc.Scan() {
		var line = sc.Text()
		cards = strings.Split(line, ": ")[1]
		winningCards = strings.Split(strings.Split(cards, "|")[0], " ")
		actualCards = strings.Split(strings.Split(cards, "|")[1], " ")
		score += intersectionPower(winningCards, actualCards)
	}
	// fmt.Println("ScratchCards: ", cards)
	// fmt.Println("winningCards: ", winningCards)
	// fmt.Println("actualCards: ", actualCards)
	fmt.Println("Score: ", score)
}

func intersectionPower(winningCards []string, actualCards []string) int {
	var instersection = intersection(winningCards, actualCards)
	var power float64 = float64(len(instersection))
	if power == 0 {
		return 0
	}
	var score int = int(math.Pow(2, power-1))
	return score
}

func intersection(winningCardsS []string, actualCardsS []string) []int {
	var instersection []int
	var winningCards []int
	var actualCards []int
	for _, winningCardS := range winningCardsS {
		winningCard, err := strconv.Atoi(winningCardS)
		if err != nil {
			fmt.Println("Error converting winningCard to int: ", err)
		} else {
			winningCards = append(winningCards, winningCard)
		}
	}
	for _, actualCardS := range actualCardsS {
		actualCard, err := strconv.Atoi(actualCardS)
		if err != nil {
			fmt.Println("Error converting actualCard to int: ", err)
		} else {
			actualCards = append(actualCards, actualCard)
		}
	}
	for _, winningCard := range winningCards {
		for _, actualCard := range actualCards {
			if winningCard == actualCard {
				instersection = append(instersection, winningCard)
			}
		}
	}
	return instersection
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
