package days

import (
	"fmt"
	"io/ioutil"
)

func getInput() string {
	filePath := "days/day1.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return string(file)
}

func day1Part1() {
	fmt.Println("day 1 part 1")
}

func day1Part2() {
	fmt.Println("day 1 part 2")
}

func Day1() {
	day1Part1()
	day1Part2()
}