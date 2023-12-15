package days

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Values struct {
	left  string
	right string
}

func parseInstructions(input string) map[string]Values {
	lines := strings.Split(input, "\n")

	m := make(map[string]Values)
	for _, line := range lines {
		split := strings.Split(line, " = ")

		key := split[0]

		values := strings.TrimLeft(strings.TrimRight(split[1], ")"), "(")

		leftAndRight := strings.Split(values, ", ")
		left := leftAndRight[0]
		right := leftAndRight[1]

		m[key] = Values{left, right}
	}

	return m
}

func getInputDay8() string {
	filePath := "days/day8.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func day8Part1(input string) {
	split := strings.Split(input, "\n\n")
	parsed := parseInstructions(split[1])

	route := split[0]
	start := "AAA"
	goal := "ZZZ"

	pointer := start
	for i := 0; i <= 100000; i++ {
		instruction := string(route[i%len(route)])

		if instruction == "L" {
			pointer = parsed[pointer].left
		}
		if instruction == "R" {
			pointer = parsed[pointer].right
		}

		if pointer == goal {
			fmt.Println(i + 1)
			break
		}
	}
}

func day8Part2(input string) {
	split := strings.Split(input, "\n\n")
	parsed := parseInstructions(split[1])

	route := split[0]

	start := []string{}
	for key := range parsed {
		if strings.HasSuffix(key, "A") {
			start = append(start, key)
		}
	}

	list := []int{}
	for _, pointer := range start {
		for i := 0; i <= 100000; i++ {
			instruction := string(route[i%len(route)])

			if instruction == "L" {
				pointer = parsed[pointer].left
			}
			if instruction == "R" {
				pointer = parsed[pointer].right
			}

			if strings.HasSuffix(pointer, "Z") {
				list = append(list, i+1)
				break
			}
		}
	}

	answer := lcm(list)
	fmt.Println(answer)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcmTwoNumbers(result, numbers[i])
	}

	return result
}

func lcmTwoNumbers(a, b int) int {
	return a * b / gcd(a, b)
}

func Day8() {
	input := getInputDay8()
	day8Part1(input)
	day8Part2(input)
}