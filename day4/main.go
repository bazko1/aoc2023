package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringArrToAtoi(arr []string) (out []int) {
	for _, el := range arr {
		val, _ := strconv.Atoi(el)
		out = append(out, val)
	}
	return out
}

func main() {
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
				if drawn == winning {
					count += 1
				}
			}
		}
		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
			// fmt.Println(sum)
		}
	}
	fmt.Println(sum)
}
