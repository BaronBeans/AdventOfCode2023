package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInputDay2() string {
	filePath := "days/day2.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func day2Part1(input string) {
	maxRed := 12
	maxBlue := 14
	maxGreen := 13

	possibleIdSum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		gameId := parts[0]
		gameInfo := parts[1]

		rounds := strings.Split(gameInfo, ";")

		var numRed, numBlue, numGreen int
		for _, round := range rounds {
			results := strings.Split(round, ",")
			for _, result := range results {
				trimmedResult := strings.TrimLeft(result, " ")
				trimmedResult = strings.TrimRight(trimmedResult, " ")
				info := strings.Split(trimmedResult, " ")

				if info[1] == "red" {
					num, _ := strconv.Atoi(info[0])
					if num > numRed {
						numRed = num
					}
				}
				if info[1] == "blue" {
					num, _ := strconv.Atoi(info[0])
					if num > numBlue {
						numBlue = num
					}
				}
				if info[1] == "green" {
					num, _ := strconv.Atoi(info[0])
					if num > numGreen {
						numGreen = num
					}
				}
			}

		}

		if numRed > maxRed {
			// fmt.Println(gameId, "red", numRed, "blue", numBlue, "green", numGreen, "impossible")
		} else if numBlue > maxBlue {
			// fmt.Println(gameId, "red", numRed, "blue", numBlue, "green", numGreen, "impossible")
		} else if numGreen > maxGreen {
			// fmt.Println(gameId, "red", numRed, "blue", numBlue, "green", numGreen, "impossible")
		} else {
			// fmt.Println(gameId, "red", numRed, "blue", numBlue, "green", numGreen, "possible")
			gameId, _ := strconv.Atoi(strings.Split(gameId, " ")[1])
			possibleIdSum = possibleIdSum + gameId
		}

	}
	fmt.Println(possibleIdSum)

}

func day2Part2(input string) {
	powerSum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		gameInfo := parts[1]

		rounds := strings.Split(gameInfo, ";")

		var numRed, numBlue, numGreen int
		for _, round := range rounds {
			results := strings.Split(round, ",")
			for _, result := range results {
				trimmedResult := strings.TrimLeft(result, " ")

				trimmedResult = strings.TrimRight(trimmedResult, " ")
				info := strings.Split(trimmedResult, " ")

				if info[1] == "red" {
					num, _ := strconv.Atoi(info[0])
					if num > numRed {
						numRed = num
					}
				}
				if info[1] == "blue" {
					num, _ := strconv.Atoi(info[0])
					if num > numBlue {
						numBlue = num
					}
				}
				if info[1] == "green" {
					num, _ := strconv.Atoi(info[0])
					if num > numGreen {
						numGreen = num
					}
				}
			}

		}

		power := numRed * numBlue * numGreen
		powerSum = powerSum + power
	}
	fmt.Println(powerSum)
}

func Day2() {
	input := getInputDay2()
	day2Part1(input)
	day2Part2(input)
}