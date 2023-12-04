package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Number struct {
	value int
	index int
}

func (n Number) ToString() string {
	return strconv.Itoa(n.value)
}
func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var score int

	for sc.Scan() {
		var currentLine = sc.Text()
		var lineScore = getScoreNumbers(currentLine)
		score += lineScore
	}
	fmt.Println("Final score:", score)
}

func getScoreNumbers(currentLine string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	values := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	fmt.Println("Current line: ", currentLine)
	// Convert currentLine to an unordered list of indexed numbers
	numberList := []Number{}
	for _, number := range numbers {
		index := strings.Index(currentLine, number)
		if index != -1 {
			partialNumberList := getPartialNumberList(currentLine, number, values)
			numberList = append(numberList, partialNumberList...)
		}
	}
	fmt.Println("Number list: ", numberList)
	sort.Slice(numberList, func(i, j int) bool {
		return numberList[i].index < numberList[j].index
	})
	fmt.Println("Sorted number list: ", numberList)
	score, err := strconv.Atoi(numberList[0].ToString() + numberList[len(numberList)-1].ToString())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Score: ", score)
	return score
}

func getPartialNumberList(currentLine string, number string, values map[string]int) []Number {
	var partialNumberList []Number
	// Number{values[number], index}
	count := strings.Count(currentLine, number)
	currentLineCopy := currentLine
	for count > 0 {
		index := strings.Index(currentLineCopy, number)
		partialNumberList = append(partialNumberList, Number{values[number], index})
		currentLineCopy = strings.Replace(currentLineCopy, number, strings.Repeat("_", len(number)), 1)
		count = strings.Count(currentLineCopy, number)
	}
	return partialNumberList
}
