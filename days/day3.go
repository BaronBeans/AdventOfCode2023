package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getInputDay3() string {
	filePath := "days/day3.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.Trim(string(file), "\n")
}

func isSymbol(r rune) bool {
	if unicode.IsDigit(r) || r == '.' {
		return false
	}

	return true
}

func isGearSymbol(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}

type Number struct {
	value       int
	stringValue string
	start       [2]int
	end         [2]int
}

func buildNumber(line string, x, y int) Number {
	result := Number{
		value:       0,
		stringValue: "0",
		start:       [2]int{0, 0},
		end:         [2]int{0, 0},
	}

	numStr := ""
	start := -1
	// left:
	for l := y; l >= 0; l-- {
		char := rune(line[l])
		if unicode.IsDigit(char) {
			numStr = string(char) + numStr
			start = l
		} else {
			break
		}
	}

	end := start
	// right
	for r := y + 1; r < len(line); r++ {
		char := rune(line[r])
		if unicode.IsDigit(char) {
			numStr = numStr + string(char)
			end = r
		} else {
			break
		}
	}

	numVal, _ := strconv.Atoi(numStr)
	result = Number{
		value:       numVal,
		stringValue: numStr,
		start:       [2]int{x, start},
		end:         [2]int{x, end},
	}

	return result

}

func alreadyExists(nums []Number, n Number) bool {
	for _, num := range nums {
		if num.value == n.value && num.start == n.start {
			return true
		}
	}
	return false
}

func day3Part1(input string) {
	lines := strings.Split(input, "\n")
	nums := []Number{}
	for row, line := range lines {
		for column, char := range line {
			if isSymbol(char) {
				for x := row - 1; x <= row+1; x++ {
					for y := column - 1; y <= column+1; y++ {
						if x < 0 || x >= len(lines) {
							continue
						}
						if y < 0 || y >= len(line) {
							continue
						}

						check := rune(lines[x][y])
						if unicode.IsDigit(check) {
							// build number:
							result := buildNumber(lines[x], x, y)
							// check if we already have the number
							if alreadyExists(nums, result) == false {
								nums = append(nums, result)
							}
						}
					}
				}
			}
		}
	}

	total := 0
	for _, num := range nums {
		total = total + num.value
	}

	fmt.Println(total)
}

func day3Part2(input string) {
	lines := strings.Split(input, "\n")
	ratios := []int{}
	for row, line := range lines {
		for column, char := range line {
			gearPartCount := 0
			gearParts := []Number{}
			if isGearSymbol(char) {
				for x := row - 1; x <= row+1; x++ {
					for y := column - 1; y <= column+1; y++ {
						if x < 0 || x >= len(lines) {
							continue
						}
						if y < 0 || y >= len(line) {
							continue
						}

						check := rune(lines[x][y])
						if unicode.IsDigit(check) {
							// build number:
							result := buildNumber(lines[x], x, y)
							// check if we already have the number
							if alreadyExists(gearParts, result) == false {
								gearPartCount = gearPartCount + 1
								gearParts = append(gearParts, result)
							}
						}
					}
				}

				if gearPartCount == 2 {
					prod := gearParts[0].value * gearParts[1].value
					ratios = append(ratios, prod)
				}
			}
		}
	}

	total := 0
	for _, ratio := range ratios {
		total = total + ratio
	}

	fmt.Println(total)
}

func Day3() {
	input := getInputDay3()
	day3Part1(input)
	day3Part2(input)
}
