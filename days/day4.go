package days

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func getInputDay4() string {
	filePath := "days/day4.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func day4Part1(input string) {
	lines := strings.Split(input, "\n")

	points := 0
	for _, line := range lines {
		split := strings.Split(line, ":")

		gameInfo := split[1]

		gameInfo = strings.ReplaceAll(gameInfo, "  ", " 0")
		gameInfoSplit := strings.Split(gameInfo, " | ")
		winning := strings.Split(strings.TrimLeft(gameInfoSplit[0], " "), " ")
		possible := strings.Split(gameInfoSplit[1], " ")

		matches := 0
		for _, number := range possible {
			for _, win := range winning {
				if string(win) == string(number) {
					matches = matches + 1
				}
			}
		}

		if matches == 0 {
			continue
		}
		pow := math.Pow(2, float64(matches-1))

		points = points + int(pow)
	}

	fmt.Println(points)
}

func processCard(cardNumber int, cardWinners, cardNumbers []string) int {
	matches := 0
	for _, number := range cardNumbers {
		for _, winner := range cardWinners {
			if string(winner) == string(number) {
				matches = matches + 1
			}
		}
	}

	return matches
}

func day4Part2(input string) {
	lines := strings.Split(input, "\n")

	totalRuns := 0
	extraCards := []int{}
	for lineNumber, line := range lines {

		runTimes := 1
		for _, i := range extraCards {
			if i == lineNumber+1 {
				runTimes = runTimes + 1
			}
		}

		totalRuns = totalRuns + runTimes

		split := strings.Split(line, ":")

		gameInfo := split[1]

		gameInfo = strings.ReplaceAll(gameInfo, "  ", " 0")
		gameInfoSplit := strings.Split(gameInfo, " | ")
		winning := strings.Split(strings.TrimLeft(gameInfoSplit[0], " "), " ")
		possible := strings.Split(gameInfoSplit[1], " ")

		numOfWins := processCard(lineNumber+1, winning, possible)
		newCards := []int{}
		for i := lineNumber + 1; i < (lineNumber+1)+numOfWins; i++ {
			newCards = append(newCards, i+1)
		}

		for i := 0; i < runTimes; i++ {
			extraCards = append(extraCards, newCards...)
		}
	}

	fmt.Println(totalRuns)
}

func Day4() {
	input := getInputDay4()
	day4Part1(input)
	day4Part2(input)
}