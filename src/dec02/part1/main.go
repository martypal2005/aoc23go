package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filePath = "../../../data/dec2/input.txt"

func getGameID(s string) int {
	parts := strings.Split(s, " ")

	if len(parts) != 2 || parts[0] != "Game" {
		return 0
	}

	gameID, err := strconv.Atoi(parts[1])
	if err == nil {
		return gameID
	}

	return 0
}

func getNumColor(s string) (int, string) {
	parts := strings.Split(s, " ")

	if len(parts) != 2 {
		return 0, ""
	}

	count, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, ""
	}

	return count, parts[1]
}

func main() {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	maxColorMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ":")

		if len(tokens) != 2 {
			fmt.Printf("Imperly formatted line: %s\n", line)
		}

		gameID := getGameID(tokens[0])

		newString := strings.ReplaceAll(tokens[1], ";", ",")

		sections := strings.Split(newString, ",")

		validGame := true
		for _, section := range sections {
			trimmed := strings.TrimSpace(section)
			count, color := getNumColor(trimmed)

			if count > maxColorMap[color] {
				validGame = false
				break
			}
		}

		if validGame {
			total += gameID
		}
	}

	fmt.Printf("--\nSum of game ids that are valid: %d\n", total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
