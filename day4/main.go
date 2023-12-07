package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func PartOne() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	var sum int
	for _, line := range lines {
		sep := strings.Split(line, "|")
		winning := sep[0]
		drawn := sep[1]
		drawn = strings.Trim(drawn, " ")
		winning = strings.Split(winning, ":")[1]
		winning = strings.Trim(winning, " ")
		winning = strings.ReplaceAll(winning, "  ", " ")
		drawn = strings.ReplaceAll(drawn, "  ", " ")
		drawnDig := strings.Split(drawn, " ")
		winningDig := strings.Split(winning, " ")
		count := 0
		for _, drawn := range drawnDig {
			for _, winning := range winningDig {
				// string comparison cuz why not
				if drawn == winning {
					count += 1
				}
			}
		}
		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
		}
	}
	fmt.Println(sum)
}

func PartTwo() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	copyCounter := make(map[int]int)
	for cardId, line := range lines {
		sep := strings.Split(line, "|")
		winning := sep[0]
		drawn := sep[1]
		drawn = strings.Trim(drawn, " ")
		winning = strings.Split(winning, ":")[1]
		winning = strings.Trim(winning, " ")
		winning = strings.ReplaceAll(winning, "  ", " ")
		drawn = strings.ReplaceAll(drawn, "  ", " ")
		drawnDig := strings.Split(drawn, " ")
		winningDig := strings.Split(winning, " ")
		count := 0
		for _, drawn := range drawnDig {
			for _, winning := range winningDig {
				// string comparison cuz why not
				if drawn == winning {
					count += 1
				}
			}
		}
		for i := cardId + 1; i < cardId+count+1; i++ {
			copyCounter[i] += 1 + copyCounter[cardId]
		}
	}
	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += copyCounter[i] + 1
	}
	fmt.Println(sum)
}

func main() {
	PartOne()
	PartTwo()
}
