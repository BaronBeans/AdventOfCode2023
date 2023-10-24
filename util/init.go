package util

import (
	"fmt"
	"io/ioutil"
)

func createDay(day int) {
	dayFilePath := fmt.Sprintf("days/day%d.go", day)
	dayFileContents := fmt.Sprintf("package days\n\nimport \"fmt\"\n\nfunc day%dPart1() {\n\tfmt.Println(\"day %d part 1\")\n}\n\nfunc day%dPart2() {\n\tfmt.Println(\"day %d part 2\")\n}\n\nfunc Day%d() {\n\tday%dPart1()\n\tday%dPart2()\n}", day, day, day, day, day, day, day)
	fmt.Println(dayFileContents)

	err := ioutil.WriteFile(dayFilePath, []byte(dayFileContents), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func Init() {
	for i := 1; i <= 30; i++ {
		createDay(i)
	}
}
