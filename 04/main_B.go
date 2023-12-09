package main

import (
	"bufio"
	"fmt"
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
	var cardLine int = 1
	var reference map[int]string = make(map[int]string)
	var solution map[int]int = make(map[int]int)
	var queue []string
	var referenceQueue []string
	for sc.Scan() {
		var line = sc.Text()
		reference[cardLine] = line
		queue = append(queue, line)
		cardLine++
	}
	cardLine = 1
	for _, line := range queue {
		cards = strings.Split(line, ": ")[1]
		winningCards = strings.Split(strings.Split(cards, "|")[0], " ")
		actualCards = strings.Split(strings.Split(cards, "|")[1], " ")
		lineScore := intersectionPower(winningCards, actualCards)
		score += lineScore
		solution[cardLine] = lineScore
		cardLine++
	}
	referenceQueue = queue
	var usedCards int = 0
	for len(queue) > 0 {
		card := queue[0]
		queue = queue[1:]
		cardId, err := strconv.Atoi(strings.Split(strings.Split(card, ": ")[0], " ")[1])
		if err != nil {
			fmt.Println("Error converting cardId to int on Id: ", err, card)
		}
		solutionIdx := solution[cardId]
		queue = append(queue, referenceQueue[cardId:cardId+solutionIdx]...)
		usedCards++
	}
	// fmt.Println("ScratchCards: ", cards)
	// fmt.Println("winningCards: ", winningCards)
	// fmt.Println("actualCards: ", actualCards)
	fmt.Println("Used cards: ", usedCards)
}

func intersectionPower(winningCards []string, actualCards []string) int {
	var instersection = intersection(winningCards, actualCards)
	// var power float64 = float64(len(instersection))
	// if power == 0 {
	// 	return 0
	// }
	// var score int = int(math.Pow(2, power-1))
	return len(instersection)
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
