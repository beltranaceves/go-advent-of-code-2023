package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var power int

	for sc.Scan() {
		var gameLine = sc.Text()
	}
	fmt.Println("Final score:", power)
}
