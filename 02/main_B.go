package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allowedRed int = 12
var allowedGreen int = 13
var allowedBlue int = 14

type Game struct {
	id    int
	pulls []Pull
}

type Pull struct {
	colors []Color
}
type Color struct {
	color string
	count int
}

type ColorRequirement struct {
	red   int
	green int
	blue  int
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g Game) GetColorRequirements() ColorRequirement {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, pull := range g.pulls {
		for _, color := range pull.colors {
			switch color.color {
			case "red":
				maxRed = MaxInt(maxRed, color.count)
			case "green":
				maxGreen = MaxInt(maxGreen, color.count)
			case "blue":
				maxBlue = MaxInt(maxBlue, color.count)
			}
		}
	}
	return ColorRequirement{maxRed, maxGreen, maxBlue}
}

func (g Game) GetPower() int {
	colorRequirements := g.GetColorRequirements()
	return colorRequirements.red * colorRequirements.green * colorRequirements.blue
}

func (g Game) Print() {
	fmt.Println("Game id:", g.id)
	for _, pull := range g.pulls {
		fmt.Println("Pull:")
		for _, color := range pull.colors {
			fmt.Println("Color:", color.color, "Count:", color.count)
		}
	}
}

func (g Game) IsValid() bool {
	for _, pull := range g.pulls {
		for _, color := range pull.colors {
			switch color.color {
			case "red":
				if color.count > allowedRed {
					return false
				}
			case "green":
				if color.count > allowedGreen {
					return false
				}
			case "blue":
				if color.count > allowedBlue {
					return false
				}
			}
		}
	}
	return true
}
func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var power int
	var games []Game

	gameIdCounter := 1
	for sc.Scan() {
		var gameLine = sc.Text()
		game := ParseGame(gameLine)
		game.id = gameIdCounter
		gameIdCounter++
		game.Print()
		games = append(games, game)
	}
	fmt.Println("Valid Games:")
	for _, game := range games {
		gamePower := game.GetPower()
		colorRequirements := game.GetColorRequirements()
		fmt.Println("Game id:", game.id, "Power:", gamePower, "Requirements:", colorRequirements)
		power += gamePower
	}
	fmt.Println("Final score:", power)
}

// Each game is parsed from a line with this format:
// Game <game_id>:<color><count>;<color><count>;...
func ParseGame(gameLine string) Game {
	fmt.Println("Game line:", gameLine)
	gameContents := strings.Split(gameLine, ":")
	gameId := strings.Split(gameContents[0], " ")[1]
	fmt.Println("Game id:", gameId)
	fmt.Println("Pulls contents:", gameContents[1])
	linePulls := strings.Split(gameContents[1], ";")
	fmt.Println("Pulls:", linePulls)
	game := Game{}
	for _, linePull := range linePulls {
		gamePull := Pull{}
		fmt.Println("Pull:", linePull)
		lineColors := strings.Split(linePull, ",")
		for _, lineColor := range lineColors {
			tmpColor := Color{}
			lineColor = strings.TrimSpace(lineColor)
			color, err := strconv.Atoi(strings.Split(lineColor, " ")[0])
			if err != nil {
				fmt.Println(err)
			}
			tmpColor.count = color
			tmpColor.color = strings.Split(lineColor, " ")[1]
			gamePull.colors = append(gamePull.colors, tmpColor)
		}
		game.pulls = append(game.pulls, gamePull)
	}
	return game
}
