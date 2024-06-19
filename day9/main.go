package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("failed to open input file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		numbers := make([]int, 0, len(fields))
		for _, n := range fields {
			num, _ := strconv.Atoi(n)
			numbers = append(numbers, num)
		}
		sum += calculate(numbers)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

func calculate(in []int) int {
	diffs := make([][]int, 0, len(in))
	diffs = append(diffs, in)
	for {
		currRow := diffs[len(diffs)-1]

		prev := currRow[0]
		allZeros := true
		newRow := make([]int, 0, len(in))
		for _, next := range currRow[1:] {
			diff := next - prev
			newRow = append(newRow, diff)
			if diff != 0 {
				allZeros = false
			}
			prev = next
		}
		if allZeros {
			break
		}
		diffs = append(diffs, newRow)
	}

	sum := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		nextRow := diffs[i]
		next := nextRow[len(nextRow)-1]
		sum += next
	}
	return sum
}
