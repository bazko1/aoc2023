package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(inputFile string) int {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	re := regexp.MustCompile("[0-9]")
	var sum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineNumbers := re.FindAllString(line, -1)
		number, err := strconv.Atoi(fmt.Sprintf("%s%s", lineNumbers[0], lineNumbers[len(lineNumbers)-1]))
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	return sum
}

func PartTwo(inputFile string) int {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")

	digitStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitReplacement := make(map[string]string)
	for id, str := range digitStrings {
		digitReplacement[str] = fmt.Sprintf("%c%d%c", str[0], id+1, str[len(str)-1])
	}
	re := regexp.MustCompile("[0-9]")
	var sum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		for key, val := range digitReplacement {
			line = strings.ReplaceAll(line, key, val)
		}
		lineNumbers := re.FindAllString(line, -1)
		number, err := strconv.Atoi(fmt.Sprintf("%s%s", lineNumbers[0], lineNumbers[len(lineNumbers)-1]))
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(PartOne("input.txt"))
	fmt.Println(PartTwo("input.txt"))
}
