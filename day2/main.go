package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	reGame := regexp.MustCompile("Game ([0-9]+)")
	reBlue := regexp.MustCompile("([0-9]+) blue")
	reGreen := regexp.MustCompile("([0-9]+) green")
	reRed := regexp.MustCompile("([0-9]+) red")
	var sum int
	redLimit := 12
	greenLimit := 13
	blueLimit := 14
	for _, line := range lines {
		if line == "" {
			continue
		}
		gameNumber, _ := strconv.Atoi(reGame.FindStringSubmatch(line)[1])
		line = strings.Split(line, ":")[1]
		ballSet := strings.Split(line, ";")
		ok := true
		for _, s := range ballSet {
			var blue, green, red int
			if blueMatch := reBlue.FindStringSubmatch(s); len(blueMatch) > 0 {
				blue, _ = strconv.Atoi(blueMatch[1])
			}
			if greenMatch := reGreen.FindStringSubmatch(s); len(greenMatch) > 0 {
				green, _ = strconv.Atoi(greenMatch[1])
			}
			if redMatch := reRed.FindStringSubmatch(s); len(redMatch) > 0 {
				red, _ = strconv.Atoi(redMatch[1])
			}
			if blue > blueLimit || green > greenLimit || red > redLimit {
				ok = false
			}
		}
		if ok {
			sum += gameNumber
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
	reBlue := regexp.MustCompile("([0-9]+) blue")
	reGreen := regexp.MustCompile("([0-9]+) green")
	reRed := regexp.MustCompile("([0-9]+) red")
	var sum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.Split(line, ":")[1]
		ballSet := strings.Split(line, ";")
		var redMax, blueMax, greenMax int
		for _, s := range ballSet {
			var blue, green, red int
			if blueMatch := reBlue.FindStringSubmatch(s); len(blueMatch) > 0 {
				blue, _ = strconv.Atoi(blueMatch[1])
			}
			if blue > blueMax {
				blueMax = blue
			}
			if greenMatch := reGreen.FindStringSubmatch(s); len(greenMatch) > 0 {
				green, _ = strconv.Atoi(greenMatch[1])
			}
			if green > greenMax {
				greenMax = green
			}
			if redMatch := reRed.FindStringSubmatch(s); len(redMatch) > 0 {
				red, _ = strconv.Atoi(redMatch[1])
			}
			if red > redMax {
				redMax = red
			}
		}
		sum += blueMax * greenMax * redMax
	}
	fmt.Println(sum)
}
func main() {
	PartOne()
	PartTwo()
}
