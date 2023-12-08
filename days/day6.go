package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInputDay6() string {
	filePath := "days/day6.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

type ParsedInput struct {
	times     []int
	distances []int
}

func parseInput(input string) ParsedInput {
	lines := strings.Split(input, "\n")
	t := strings.Split(lines[0], " ")
	d := strings.Split(lines[1], " ")
	times := []int{}
	for _, time := range t {
		i, err := strconv.Atoi(time)
		if err != nil {
			continue
		}
		times = append(times, i)
	}
	distances := []int{}
	for _, distance := range d {
		i, err := strconv.Atoi(distance)
		if err != nil {
			continue
		}
		distances = append(distances, i)
	}
	return ParsedInput{
		times:     times,
		distances: distances,
	}
}

func parseInputPart2(input string) []int {
	lines := strings.Split(input, "\n")
	t := strings.Split(lines[0], " ")[1:]
	d := strings.Split(lines[1], " ")[1:]
	times := ""
	for _, time := range t {
		times = times + time
	}
	distances := ""
	for _, distance := range d {
		distances = distances + distance
	}

	time, err := strconv.Atoi(times)
	if err != nil {
		fmt.Println("error parsing time")
	}
	distance, err := strconv.Atoi(distances)
	if err != nil {
		fmt.Println("error parsing time")
	}

	return []int{time, distance}
}

func day6Part1(input string) {
	parsed := parseInput(input)

	finalAnswer := 1
	for r := 0; r < len(parsed.times); r++ {
		time := parsed.times[r]
		distance := parsed.distances[r]

		possibilityCount := 0
		for i := 0; i < time; i++ {
			result := (time - i) * i
			if result > distance {
				possibilityCount = possibilityCount + 1
			}
		}
		finalAnswer = finalAnswer * possibilityCount
	}
	fmt.Println(finalAnswer)
}

func day6Part2(input string) {
	parsed := parseInputPart2(input)
	time := parsed[0]
	distance := parsed[1]

	possibilityCount := 0
	for i := 0; i < time; i++ {
		result := (time - i) * i
		if result > distance {
			possibilityCount = possibilityCount + 1
		}
	}

	fmt.Println(possibilityCount)
}

func Day6() {
	input := getInputDay6()
	day6Part1(input)
	day6Part2(input)
}