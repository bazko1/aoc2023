package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
				}
				digitStack = ""
				isAdjacent = false
			}
		}
	}
	fmt.Println(sum)
}

type Point struct {
	y, x int
}

func PartTwo() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	var sum int
	digitStack := ""
	adjacentDigitsStack := make([]Point, 0)
	adjacentDigitsToValue := make(map[Point]int)
	countedStar := make(map[Point]bool)
	isAdjacent := false
	adjacentStar := make(map[Point][]Point)
	for row, line := range lines {
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				digitStack += string(ch)
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if X, Y := col+x, row+y; X >= 0 && X < len(line) &&
							Y >= 0 && Y < len(lines) &&
							[]rune(lines[Y])[X] == '*' {
							adjacentDigitsStack = append(adjacentDigitsStack, Point{row, col})
							isAdjacent = true
							if _, counted := countedStar[Point{Y, X}]; !counted {
								if _, has := adjacentStar[Point{Y, X}]; !has {
									adjacentStar[Point{Y, X}] = make([]Point, 0)
								}
								adjacentStar[Point{Y, X}] = append(adjacentStar[Point{Y, X}], Point{row, col})
								countedStar[Point{Y, X}] = true
							}
						}
					}
				}
			} else {
				if digitStack != "" && isAdjacent {
					val, _ := strconv.Atoi(digitStack)
					for _, p := range adjacentDigitsStack {
						adjacentDigitsToValue[p] = val
					}
				}
				digitStack = ""
				adjacentDigitsStack = make([]Point, 0)
				countedStar = make(map[Point]bool)
				isAdjacent = false
			}
		}
	}
	for _, arr := range adjacentStar {
		if len(arr) == 2 {
			sum += adjacentDigitsToValue[arr[0]] * adjacentDigitsToValue[arr[1]]
		}
	}
	fmt.Println(sum)
}

func main() {
	PartOne()
	PartTwo()
}
