package days

import (
	"fmt"
	"os"
	"strings"
)

func getInputDay10() string {
	filePath := "days/day10.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func day10Part1(input string) {
	fmt.Println("day 10 part 1")
}

func day10Part2(input string) {
	fmt.Println("day 10 part 2")
}

func Day10() {
	input := getInputDay10()
	day10Part1(input)
	day10Part2(input)
}
