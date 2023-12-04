package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var score int

	for sc.Scan() {
		var currentLine = sc.Text()
		var lineScore = getScore(currentLine)
		score += lineScore
	}
	fmt.Println(score)
}

func getScore(currentLine string) int {
	var scoreDigits string
	for _, a := range currentLine {
		if unicode.IsNumber(a) {
			fmt.Println("Number: ", string(a))
			scoreDigits += string(a)
		}
	}

	scoreChars := strings.Split(scoreDigits, "")
	scoreList := []string{scoreChars[0], scoreChars[len(scoreChars)-1]}
	scorePoints := strings.Join(scoreList, "")
	score, err := strconv.Atoi(scorePoints)
	if err != nil {
		fmt.Println(err)
	}
	return score
}
