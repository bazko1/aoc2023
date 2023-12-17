package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func AtoiArr(arr []string) (o []float64) {
	for _, el := range arr {
		val, _ := strconv.Atoi(el)
		o = append(o, float64(val))
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
	timeLine := lines[0]
	distanceLine := lines[1]
	reDigit := regexp.MustCompile("[0-9]+")
	timeStrings := reDigit.FindAllString(timeLine, -1)
	times := AtoiArr(timeStrings)
	distanceStrings := reDigit.FindAllString(distanceLine, -1)
	distances := AtoiArr(distanceStrings)
	partTwoTime, _ := strconv.ParseFloat(strings.Join(timeStrings, ""), 64)
	partTwoDist, _ := strconv.ParseFloat(strings.Join(distanceStrings, ""), 64)
	result := 1.
	// what we need is integer bigest and lowest solution of:
	// -speed**2 + speed*time - record > 0
	for i := range times {
		time := times[i]
		record := distances[i]
		dSqrt := math.Sqrt(time*time - 4*record)
		lowerBound := math.Floor((time-dSqrt)/2) + 1
		upperBound := math.Ceil((time+dSqrt)/2) - 1
		result *= (upperBound - lowerBound + 1)
	}
	fmt.Println("Part1:", result)
	dSqrtDivided := math.Sqrt(partTwoTime*partTwoTime/4 - partTwoDist)
	lowerBound := math.Floor((partTwoTime/2 - dSqrtDivided)) + 1
	upperBound := math.Ceil((partTwoTime/2 + dSqrtDivided)) - 1
	fmt.Println("Part2:", int64(upperBound-lowerBound+1))
}
