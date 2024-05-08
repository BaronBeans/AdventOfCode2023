package days

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInputDay5() string {
	filePath := "days/day5.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func mapValue(mapData []string, value int) int {
	for _, line := range mapData {
		parts := strings.Split(line, " ")
		newVal, _ := strconv.Atoi(parts[0])
		startVal, _ := strconv.Atoi(parts[1])
		lengthVal, _ := strconv.Atoi(parts[2])

		destStart := newVal
		sourceStart := startVal
		sourceEnd := startVal + (lengthVal - 1)

		diff := destStart - sourceStart

		if value >= sourceStart && value <= sourceEnd {
			return value + diff
		}
	}

	return value
}

func reverseMapValue(mapData []string, value int) int {
	for _, line := range mapData {
		parts := strings.Split(line, " ")
		newVal, _ := strconv.Atoi(parts[0])
		startVal, _ := strconv.Atoi(parts[1])
		lengthVal, _ := strconv.Atoi(parts[2])

		destStart := newVal
		destEnd := newVal + (lengthVal - 1)
		sourceStart := startVal

		diff := sourceStart - destStart

		if value >= destStart && value <= destEnd {
			newValue := value + diff
			return newValue
		}
	}

	return value
}

func mapValueFunc(value int, lines []string, reverse bool) int {
	seedToSoil := strings.Split(lines[1], "\n")[1:]
	soilToFert := strings.Split(lines[2], "\n")[1:]
	fertToWater := strings.Split(lines[3], "\n")[1:]
	waterToLight := strings.Split(lines[4], "\n")[1:]
	lightToTemp := strings.Split(lines[5], "\n")[1:]
	tempToHum := strings.Split(lines[6], "\n")[1:]
	humToLoc := strings.Split(lines[7], "\n")[1:]

	if reverse == true {
		hum := reverseMapValue(humToLoc, value)
		temp := reverseMapValue(tempToHum, hum)
		light := reverseMapValue(lightToTemp, temp)
		water := reverseMapValue(waterToLight, light)
		fert := reverseMapValue(fertToWater, water)
		soil := reverseMapValue(soilToFert, fert)
		seed := reverseMapValue(seedToSoil, soil)

		return seed
	}

	soil := mapValue(seedToSoil, value)
	fert := mapValue(soilToFert, soil)
	water := mapValue(fertToWater, fert)
	light := mapValue(waterToLight, water)
	temp := mapValue(lightToTemp, light)
	hum := mapValue(tempToHum, temp)
	loc := mapValue(humToLoc, hum)

	return loc
}

func day5Part1(input string) {
	splits := strings.Split(input, "\n\n")

	seeds := strings.Split(strings.Trim(strings.Split(splits[0], ":")[1], " "), " ")

	lowestLocation := 0
	for _, seed := range seeds {
		seed, _ := strconv.Atoi(seed)

		loc := mapValueFunc(seed, splits, false)

		if lowestLocation == 0 {
			lowestLocation = loc
		} else {
			lowestLocation = int(math.Min(float64(lowestLocation), float64(loc)))
		}
	}

	fmt.Println(lowestLocation)
}

type Range struct {
	start int
	end   int
}

func day5Part2(input string) {
	splits := strings.Split(input, "\n\n")

	seeds := strings.Split(strings.Trim(strings.Split(splits[0], ":")[1], " "), " ")

	ranges := []Range{}
	for s := 0; s < len(seeds); s = s + 2 {
		seedStartVal, _ := strconv.Atoi(seeds[s])
		seedRangeVal, _ := strconv.Atoi(seeds[s+1])
		seedStart := seedStartVal
		seedEnd := seedStartVal + (seedRangeVal - 1)

		ranges = append(ranges, Range{
			start: seedStart,
			end:   seedEnd,
		})
	}

outerLoop:
	for i := 2_500_000; i < 2_600_000; i++ {
		seed := mapValueFunc(i, splits, true)

		for _, r := range ranges {
			if seed >= r.start && seed <= r.end {
				fmt.Println(i)
				break outerLoop
			}
		}
	}
}

func Day5() {
	input := getInputDay5()
	day5Part1(input)
	day5Part2(input)
}
