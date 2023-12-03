package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	bytes, err := os.ReadFile("simple.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	digitRe := regexp.MustCompile("[0-9]+")
	symbolRe := regexp.MustCompile("[^.]+")
	var symbols [][]int
	var digits [][]int
	for _, line := range lines {
		lineDigits := digitRe.FindAllStringIndex(line, -1)
		lineSymbols := symbolRe.FindAllStringIndex(line, -1)
		fmt.Println(lineDigits)
		fmt.Println("symb:=", lineSymbols)
	}
}
