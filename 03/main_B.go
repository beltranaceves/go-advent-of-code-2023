package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PartNumber struct {
	partNumber       string
	startCoordinates Coordinates
	endCoordinates   Coordinates
}

type Coordinates struct {
	row    int
	column int
}

func (p PartNumber) Check(engineMatrix [][]string) (bool, [][]string) {
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "."}
	checkAdjacent := func(row int, column int, engineMatrix [][]string) bool {
		if row > 0 {
			if !slices.Contains(numbers, engineMatrix[row-1][column]) {
				return true
			}
		}
		if row < len(engineMatrix)-1 {
			if !slices.Contains(numbers, engineMatrix[row+1][column]) {
				return true
			}
		}
		if column > 0 {
			if !slices.Contains(numbers, engineMatrix[row][column-1]) {
				return true
			}
		}
		if column < len(engineMatrix[row])-1 {
			if !slices.Contains(numbers, engineMatrix[row][column+1]) {
				return true
			}
		}
		if row > 0 && column > 0 {
			if !slices.Contains(numbers, engineMatrix[row-1][column-1]) {
				return true
			}
		}
		if row > 0 && column < len(engineMatrix[row])-1 {
			if !slices.Contains(numbers, engineMatrix[row-1][column+1]) {
				return true
			}
		}
		if row < len(engineMatrix)-1 && column > 0 {
			if !slices.Contains(numbers, engineMatrix[row+1][column-1]) {
				return true
			}
		}
		if row < len(engineMatrix)-1 && column < len(engineMatrix[row])-1 {
			if !slices.Contains(numbers, engineMatrix[row+1][column+1]) {
				return true
			}
		}
		return false
	}

	for row := p.startCoordinates.row; row <= p.endCoordinates.row; row++ {
		for column := p.startCoordinates.column; column <= p.endCoordinates.column; column++ {
			// fmt.Println("Checking", row, column, engineMatrix[row][column])
			if checkAdjacent(row, column, engineMatrix) {
				return true, engineMatrix
			} else {
				engineMatrix[row][column] = "."
			}
		}
	}
	return false, engineMatrix
}

func (p PartNumber) GetNumber() int {
	number, _ := strconv.Atoi(p.partNumber)
	return number
}

type Gear struct {
	coordinates Coordinates
}

func (g Gear) GetAdjacentPartNumbers(partNumbers []PartNumber) []PartNumber {
	var adjacentPartNumbers []PartNumber
	valuesToCheck := []Coordinates{
		Coordinates{g.coordinates.row - 1, g.coordinates.column},
		Coordinates{g.coordinates.row + 1, g.coordinates.column},
		Coordinates{g.coordinates.row, g.coordinates.column - 1},
		Coordinates{g.coordinates.row, g.coordinates.column + 1},
		Coordinates{g.coordinates.row - 1, g.coordinates.column - 1},
		Coordinates{g.coordinates.row - 1, g.coordinates.column + 1},
		Coordinates{g.coordinates.row + 1, g.coordinates.column - 1},
		Coordinates{g.coordinates.row + 1, g.coordinates.column + 1},
	}
	// fmt.Println("Checking adjacent part numbers for gear", g.coordinates)
	// fmt.Println("Values to check:", valuesToCheck)
	for _, coordinate := range valuesToCheck {
		// fmt.Println("Checking coordinate", coordinate)
		for _, partNumber := range partNumbers {
			// fmt.Println("Checking part number", partNumber)
			if partNumber.startCoordinates.row <= coordinate.row && partNumber.endCoordinates.row >= coordinate.row {
				if partNumber.startCoordinates.column <= coordinate.column && partNumber.endCoordinates.column >= coordinate.column {
					// fmt.Println("Gear", g.coordinates, "is in part number", partNumber.partNumber)
					if !slices.Contains(adjacentPartNumbers, partNumber) {
						adjacentPartNumbers = append(adjacentPartNumbers, partNumber)
					}
				}
			}
		}
	}
	return adjacentPartNumbers
}

func main() {
	//Read input file
	// fmt.Println("Args:", os.Args)
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var engineMatrix [][]string
	var partNumbers []PartNumber
	var gears []Gear
	for sc.Scan() {
		var engineRow = sc.Text()
		engineMatrix = append(engineMatrix, strings.Split(engineRow, ""))
		partNumbers = ParsePartNumbers(engineMatrix)
		gears = ParseGears(engineMatrix)
	}
	partNumbers = slices.DeleteFunc(partNumbers, func(p PartNumber) bool {
		val, _ := p.Check(engineMatrix)
		return !val
	})

	gearRatioSum := 0

	for _, gear := range gears {
		// fmt.Println("Gear:", gear)
		adjacentPartNumbers := gear.GetAdjacentPartNumbers(partNumbers)
		fmt.Println("Adjacent part numbers:", adjacentPartNumbers)
		if len(adjacentPartNumbers) == 2 {
			fmt.Println("Gear", gear.coordinates, "has 2 adjacent part numbers")
			gearRatioSum += adjacentPartNumbers[0].GetNumber() * adjacentPartNumbers[1].GetNumber()
		}
	}

	// fmt.Println("Engine:", engineMatrix)
	// fmt.Println("Part numbers:", partNumbers)
	fmt.Println("Total Gear ratio:", gearRatioSum)
}

func ParseGears(engineMatrix [][]string) []Gear {
	var gears []Gear
	for rowIdx, row := range engineMatrix {
		for columnIdx, columnNumber := range row {
			if columnNumber == "*" {
				gear := Gear{
					coordinates: Coordinates{rowIdx, columnIdx},
				}
				gears = append(gears, gear)
			}
		}
	}
	return gears
}

func ParsePartNumbers(engineMatrix [][]string) []PartNumber {
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	// separator := "."
	var partNumbers []PartNumber
	var numberBuffer strings.Builder
	var startCoordinates Coordinates
	var endCoordinates Coordinates
	var lastCoordinates Coordinates
	creatingPartNumber := false
	for rowIdx, row := range engineMatrix {
		for columnIdx, columnNumber := range row {
			if slices.Contains(numbers, columnNumber) {
				if !creatingPartNumber {
					creatingPartNumber = true
					startCoordinates = Coordinates{rowIdx, columnIdx}
				}
				lastCoordinates = Coordinates{rowIdx, columnIdx}
				numberBuffer.WriteString(columnNumber)
			} else {
				if creatingPartNumber {
					creatingPartNumber = false
					endCoordinates = lastCoordinates
					partNumber := PartNumber{
						partNumber:       numberBuffer.String(),
						startCoordinates: startCoordinates,
						endCoordinates:   endCoordinates,
					}
					partNumbers = append(partNumbers, partNumber)
					numberBuffer.Reset()
				}
			}
		}
	}
	return partNumbers
}
