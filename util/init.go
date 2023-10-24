package util

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func updateMain(day int) {
	mainFilePath := "main.go"
	mainContents, errRead := ioutil.ReadFile(mainFilePath)

	if errRead != nil {
		fmt.Println("Error reading file:", errRead)
		return
	}

	newMainContents := strings.Replace(string(mainContents), "//insert here", fmt.Sprintf("\"day%d\": days.Day%d\n\t//insert here", day, day), 1)

	errWrite := ioutil.WriteFile(mainFilePath, []byte(newMainContents), 0644)

	if errWrite != nil {
		fmt.Println("Error writing file:", errWrite)
		return
	}

	fmt.Println("File written successfully")
}

func Init(day int) {
	createDay(day)
	updateMain(day)
}
