package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	var sum int
	digitStack := ""
	isAdjacent := false
	for row, line := range lines {
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				digitStack += string(ch)
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if X, Y := col+x, row+y; X >= 0 && X < len(line) &&
							Y >= 0 && Y < len(lines) &&
							!unicode.IsDigit([]rune(lines[Y])[X]) &&
							[]rune(lines[Y])[X] != '.' {
							isAdjacent = true
						}
					}
				}
			} else {
				if digitStack != "" && isAdjacent {
					val, _ := strconv.Atoi(digitStack)
					sum += val
					// fmt.Println("Adding", val, col, row)
				}
				digitStack = ""
				isAdjacent = false
			}
		}
	}
	fmt.Println(sum)
}
