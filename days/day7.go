package days

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	originalHand string
	bid          int
	cardValues   []int
	cardCounts   map[int]int
	cardScore    int
	gameScore    int
	containsJs   bool
}

func (h Hand) getGameScorePart1() int {
	if len(h.cardCounts) == 1 {
		// five of a kind
		return 6
	}

	if len(h.cardCounts) == 2 {
		for _, v := range h.cardCounts {
			if v == 4 {
				// four of a kind
				return 5
			}
		}
		// full house
		return 4
	}

	if len(h.cardCounts) == 3 {
		for _, v := range h.cardCounts {
			if v == 3 {
				// three of a kind
				return 3
			}
		}
		// two pair
		return 2
	}

	if len(h.cardCounts) == 4 {
		// one pair
		return 1
	}

	return 0
}

func (h Hand) getGameScorePart2() int {
	if h.containsJs == true {
		if len(h.cardCounts) == 1 {
			// five of a kind
			return 6
		}

		if len(h.cardCounts) == 2 {
			// five of a kind
			return 6
		}

		if len(h.cardCounts) == 3 {
			if h.cardCounts[1] >= 2 {
				// four of a kind
				return 5
			}
			for _, v := range h.cardCounts {
				if v == 3 {
					// four of a kind
					return 5
				}
			}
			// full house
			return 4
		}

		if len(h.cardCounts) == 4 {
			// three of a kind
			return 3
		}

		if len(h.cardCounts) == 5 {
			return 1
		}
	} else {
		if len(h.cardCounts) == 1 {
			// five of a kind
			return 6
		}

		if len(h.cardCounts) == 2 {
			for _, v := range h.cardCounts {
				if v == 4 {
					// four of a kind
					return 5
				}
			}
			// full house
			return 4
		}

		if len(h.cardCounts) == 3 {
			for _, v := range h.cardCounts {
				if v == 3 {
					// three of a kind
					return 3
				}
			}
			// two pair
			return 2
		}

		if len(h.cardCounts) == 4 {
			// one pair
			return 1
		}
	}
	return 0
}

type ByGameScoreThenScore []Hand

func (a ByGameScoreThenScore) Len() int      { return len(a) }
func (a ByGameScoreThenScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByGameScoreThenScore) Less(i, j int) bool {
	if a[i].gameScore == a[j].gameScore {
		return a[i].cardScore < a[j].cardScore
	}
	return a[i].gameScore < a[j].gameScore
}

func parseCardsPart1(input string) []Hand {
	lines := strings.Split(input, "\n")

	dict := make(map[string]string)

	dict["2"] = "02"
	dict["3"] = "03"
	dict["4"] = "04"
	dict["5"] = "05"
	dict["6"] = "06"
	dict["7"] = "07"
	dict["8"] = "08"
	dict["9"] = "09"
	dict["T"] = "10"
	dict["J"] = "11"
	dict["Q"] = "12"
	dict["K"] = "13"
	dict["A"] = "14"

	hands := []Hand{}
	for _, line := range lines {

		split := strings.Split(line, " ")
		originalHand := split[0]
		bid := split[1]

		m := make(map[int]int)
		c := []int{}
		s := ""
		for _, char := range strings.Split(line, " ")[0] {
			val := dict[string(char)]
			valInt, _ := strconv.Atoi(val)
			m[valInt]++
			s = s + val
			c = append(c, valInt)
		}
		score, _ := strconv.Atoi(s)
		bidVal, _ := strconv.Atoi(bid)
		hand := Hand{
			originalHand: originalHand,
			bid:          bidVal,
			cardValues:   c,
			cardCounts:   m,
			cardScore:    score,
			containsJs:   false,
		}

		gameScore := hand.getGameScorePart1()
		hand.gameScore = gameScore
		hands = append(hands, hand)
	}

	return hands
}

func parseCardsPart2(input string) []Hand {
	lines := strings.Split(input, "\n")

	dict := make(map[string]string)

	dict["J"] = "01"
	dict["2"] = "02"
	dict["3"] = "03"
	dict["4"] = "04"
	dict["5"] = "05"
	dict["6"] = "06"
	dict["7"] = "07"
	dict["8"] = "08"
	dict["9"] = "09"
	dict["T"] = "10"
	dict["Q"] = "12"
	dict["K"] = "13"
	dict["A"] = "14"

	hands := []Hand{}
	for _, line := range lines {

		split := strings.Split(line, " ")
		originalHand := split[0]
		bid := split[1]

		m := make(map[int]int)
		c := []int{}
		s := ""
		containsJs := false
		for _, char := range strings.Split(line, " ")[0] {
			if string(char) == "J" {
				containsJs = true
			}
			val := dict[string(char)]
			valInt, _ := strconv.Atoi(val)
			m[valInt]++
			s = s + val
			c = append(c, valInt)
		}
		score, _ := strconv.Atoi(s)
		bidVal, _ := strconv.Atoi(bid)
		hand := Hand{
			originalHand: originalHand,
			bid:          bidVal,
			cardValues:   c,
			cardCounts:   m,
			cardScore:    score,
			containsJs:   containsJs,
		}

		gameScore := hand.getGameScorePart2()
		hand.gameScore = gameScore
		hands = append(hands, hand)
	}

	return hands
}

func getInputDay7() string {
	filePath := "days/day7.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func day7Part1(input string) {
	hands := parseCardsPart1(input)
	sort.Sort(ByGameScoreThenScore(hands))
	total := 0
	for i, hand := range hands {
		total = total + (i+1)*hand.bid
	}

	fmt.Println(total)
}

func day7Part2(input string) {
	hands := parseCardsPart2(input)
	sort.Sort(ByGameScoreThenScore(hands))
	total := 0
	for i, hand := range hands {
		total = total + (i+1)*hand.bid
	}

	fmt.Println(total)
}

func Day7() {
	input := getInputDay7()
	day7Part1(input)
	day7Part2(input)
}
