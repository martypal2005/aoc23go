package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filePath = "../../../data/dec4/input.txt"

type gameStruct struct {
	id             int
	score          int
	winningNumbers map[int]bool
	pickedNumbers  map[int]bool
}

func getGame(line string) (gameStruct, bool) {
	parts := strings.Split(line, ":")
	game := gameStruct{
		winningNumbers: make(map[int]bool),
		pickedNumbers:  make(map[int]bool),
	}
	found := false

	// First find the id
	if len(parts) != 2 {
		fmt.Printf("Invalid line: %s [SKIP]\n", line)
		return game, found
	}

	re := regexp.MustCompile(`\d+`)
	numberStr := re.FindString(line)
	if numberStr == "" {
		fmt.Printf("No number found in string %s [SKIP]\n", line)
		return game, found
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Printf("Number parsing error %s [SKIP]\n", line)
		return game, found
	}

	game.id = number

	// Get the winning numbers on the left
	numberParts := strings.Split(parts[1], "|")

	if len(numberParts) != 2 {
		fmt.Printf("Invalid number of parts in game %s [SKIP]\n", line)
	}

	winningNumbersStr := re.FindAllString(numberParts[0], -1)

	for _, winningNumber := range winningNumbersStr {
		number, err := strconv.Atoi(winningNumber)
		if err != nil {
			continue
		}

		game.winningNumbers[number] = true
	}

	// Get the picked numbers on the right
	pickedNumbersStr := re.FindAllString(numberParts[1], -1)

	for _, pickedNumbers := range pickedNumbersStr {
		number, err := strconv.Atoi(pickedNumbers)
		if err != nil {
			continue
		}

		game.pickedNumbers[number] = true
	}

	// Calculate the score
	for number := range game.pickedNumbers {
		_, exists := game.winningNumbers[number]
		if exists {
			if game.score == 0 {
				game.score = 1
			} else {
				game.score *= 2
			}
		}
	}

	return game, true
}

func main() {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		game, found := getGame(line)
		if !found {
			continue
		}
		total += game.score
	}

	fmt.Println(total)
}
