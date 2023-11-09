package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func createDayFromTemplate(day int) {
	dayFilePath := fmt.Sprintf("days/day%d.go", day)
	file, err := os.Create(dayFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()
	tmplString := template.Must(template.ParseFiles("day.template"))
	err = tmplString.Execute(file, day)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func updateMain(day int) {
	mainFilePath := "main.go"
	mainContents, errRead := ioutil.ReadFile(mainFilePath)

	if errRead != nil {
		fmt.Println("Error reading file:", errRead)
		return
	}

	newMainContents := strings.Replace(string(mainContents), "//insert here", fmt.Sprintf("\"day%d\": days.Day%d,\n\t//insert here", day, day), 1)

	errWrite := ioutil.WriteFile(mainFilePath, []byte(newMainContents), 0644)

	if errWrite != nil {
		fmt.Println("Error writing file:", errWrite)
		return
	}
}

func Init(day int) {
	fmt.Println(day)
	createDayFromTemplate(day)
	updateMain(day)
}
