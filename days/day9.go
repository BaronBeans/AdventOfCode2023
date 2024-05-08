package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputDay9() string {
	filePath := "days/day9.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func transform(nums []int) []int {
	diffs := []int{}
	for x := 0; x < len(nums)-1; x++ {
		diff := nums[x+1] - nums[x]
		diffs = append(diffs, diff)
	}
	return diffs
}

func checkNums(nums []int) bool {
	containsAllZeroes := true
	for _, num := range nums {
		if num != 0 {
			containsAllZeroes = false
		}
	}

	return containsAllZeroes
}

func day9Part1(input string) {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		nums := []int{}

		for _, str := range strings.Split(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string '%s' to int: %v\n", str, err)
				return
			}
			nums = append(nums, num)
		}

		allNums := [][]int{nums}
		for i := 0; i <= 10000; i++ {
			if checkNums(allNums[len(allNums)-1]) == false {
				newNums := transform(allNums[len(allNums)-1])
				allNums = append(allNums, newNums)
			}
		}

		for i := len(allNums) - 2; i >= 0; i-- {
			lastItem := allNums[i][len(allNums[i])-1]

			addItem := allNums[i+1][len(allNums[i+1])-1]

			allNums[i] = append(allNums[i], lastItem+addItem)
		}

		total = total + allNums[0][len(allNums[0])-1]
	}
	fmt.Println(total)
}

func day9Part2(input string) {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		nums := []int{}

		for _, str := range strings.Split(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string '%s' to int: %v\n", str, err)
				return
			}
			nums = append(nums, num)
		}

		allNums := [][]int{nums}
		for i := 0; i <= 10000; i++ {
			if checkNums(allNums[len(allNums)-1]) == false {
				newNums := transform(allNums[len(allNums)-1])
				allNums = append(allNums, newNums)
			}
		}

		for i := len(allNums) - 2; i >= 0; i-- {
			lastItem := allNums[i][0]

			addItem := allNums[i+1][0]

			allNums[i] = append([]int{lastItem - addItem}, allNums[i]...)
		}

		total = total + allNums[0][0]
	}
	fmt.Println(total)
}

func Day9() {
	input := getInputDay9()
	day9Part1(input)
	day9Part2(input)
}
