package days

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func getInput() string {
	filePath := "days/day1.input"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return string(file)
}

func day1Part1(input string) {
	allNumbers := []int{}
	output := 0
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		numbers := []int{}
		for _, char := range []rune(line) {
			if unicode.IsNumber(char) {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					fmt.Println("error converting number")
				}
				numbers = append(numbers, num)
			}
		}
		if len(numbers) > 0 {
			both := numbers[0]*10 + numbers[len(numbers)-1]
			allNumbers = append(allNumbers, both)
			output = output + both
		}
	}
	fmt.Println(output)
}

func day1Part2(input string) {
	dict := make(map[string]string)

	dict["one"] = "1"
	dict["two"] = "2"
	dict["three"] = "3"
	dict["four"] = "4"
	dict["five"] = "5"
	dict["six"] = "6"
	dict["seven"] = "7"
	dict["eight"] = "8"
	dict["nine"] = "9"

	newInput := ""
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var order []struct {
			Key      string
			Position int
		}

		for key := range dict {
			if strings.Contains(line, key) {
				order = append(order, struct {
					Key      string
					Position int
				}{key, strings.Index(line, key)})
			}
		}
		sort.Slice(order, func(i, j int) bool {
			return order[i].Position < order[j].Position
		})

		for _, value := range order {
			key := value.Key
			replace := dict[key] + "" + key

			line = strings.ReplaceAll(line, key, replace)
		}
		newInput = newInput + line + "\n"
	}

	day1Part1(newInput)
}

func Day1() {
	input := getInput()
	day1Part1(input)
	day1Part2(input)
}
