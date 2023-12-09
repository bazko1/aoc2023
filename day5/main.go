package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func AtoiArr(arr []string) (o []int) {
	for _, el := range arr {
		val, _ := strconv.Atoi(el)
		o = append(o, val)
	}
	return o
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	numberRegex := regexp.MustCompile("[0-9]+")
	seeds := AtoiArr(numberRegex.FindAllString(lines[0], -1))
	ranges := [][]int{}
	for _, line := range lines[2:] {
		if line == "" {
			for sId, s := range seeds {
				for _, r := range ranges {
					if s > r[1] && s < r[1]+r[2] {
						seeds[sId] = r[0] + s - r[1]
					}
				}
			}
			ranges = [][]int{}
		} else if line[len(line)-1] == ':' {
			continue
		} else {
			rangeData := AtoiArr(numberRegex.FindAllString(line, -1))
			if len(rangeData) > 0 {
				ranges = append(ranges, rangeData)
			}
		}
	}
	min := seeds[0]
	for _, s := range seeds {
		if min > s {
			min = s
		}
	}
	fmt.Println(min)
}
