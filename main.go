package main

import (
	"aoc-go/days"
	"aoc-go/util"
	"flag"
	"fmt"
	"os"
)

func main() {
	init := flag.Bool("i", false, "initialise new day")
	day := flag.Int("d", 0, "provide the day to use")
	flag.Parse()

	// error if there are no flags
	if *day == 0 && *init == false {
		fmt.Println("Please provide at least one flag")
		os.Exit(0)
	}

	// main logic
	if *init == true {
		util.Init()
	} else {
		runOne(day)
	}
}

func runOne(day *int) {
	str := fmt.Sprintf("day%d", *day)
	funcMap[str].(func())()
}

var funcMap = map[string]interface{}{
	"day1":  days.Day1,
	"day2":  days.Day2,
	"day3":  days.Day3,
	"day4":  days.Day4,
	"day5":  days.Day5,
	"day6":  days.Day6,
	"day7":  days.Day7,
	"day8":  days.Day8,
	"day9":  days.Day9,
	"day10": days.Day10,
	"day11": days.Day11,
	"day12": days.Day12,
	"day13": days.Day13,
	"day14": days.Day14,
	"day15": days.Day15,
	"day16": days.Day16,
	"day17": days.Day17,
	"day18": days.Day18,
	"day19": days.Day19,
	"day20": days.Day20,
	"day21": days.Day21,
	"day22": days.Day22,
	"day23": days.Day23,
	"day24": days.Day24,
	"day25": days.Day25,
	"day26": days.Day26,
	"day27": days.Day27,
	"day28": days.Day28,
	"day29": days.Day29,
	"day30": days.Day30,
}
