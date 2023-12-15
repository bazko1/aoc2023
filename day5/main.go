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

func PartOne() {
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
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Interval struct {
	start, length int
}

// Intersection of i1 and i2 and rest of i1
// meaning i1 - intersection
// return is of form {intersection, left removal part, right removal part}
func IntersectAndRest(i1, i2 Interval) []Interval {
	i1End := i1.start + i1.length
	i2End := i2.start + i2.length
	if i2.start >= i1End || i1.start >= i2End {
		return []Interval{{}, {}, i1}
	}
	newStart := Max(i1.start, i2.start)
	newEnd := Min(i1End, i2End)
	out := []Interval{{newStart, newEnd - newStart}, {}, {}}
	if newStart > i1.start {
		out[1] = Interval{i1.start, newStart - i1.start}
	}
	if newEnd < i1End {
		out[2] = Interval{newEnd, i1End - newEnd}
	}
	return out
}

func PartTwo() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	numberRegex := regexp.MustCompile("[0-9]+")
	seedRanges := []Interval{}
	seeds := AtoiArr(numberRegex.FindAllString(lines[0], -1))
	for sId := 0; sId < len(seeds)-1; sId += 2 {
		seedRanges = append(seedRanges, Interval{seeds[sId], seeds[sId+1]})
	}
	ranges := [][]int{}
	for _, line := range lines[2:] {
		if line == "" {
			queue := make([]Interval, len(seedRanges))
			for qId, q := range seedRanges {
				queue[qId] = q
			}
			seedRanges = []Interval{}
			for len(queue) > 0 {
				s := queue[0]
				queue = queue[1:]
				hadIntersection := false
				for _, r := range ranges {
					rDest := r[0]
					rSource := r[1]
					rRange := r[2]
					interAndRest := IntersectAndRest(s, Interval{rSource, rRange})
					intersect := interAndRest[0]
					leftRest := interAndRest[1]
					rightRest := interAndRest[2]
					if intersect.length > 0 {
						hadIntersection = true
						seedRanges = append(seedRanges, Interval{rDest + intersect.start - rSource, intersect.length}) // TODO: Mapping
						if leftRest.length > 0 {
							queue = append(queue, leftRest)
						}
						if rightRest.length > 0 {
							queue = append(queue, rightRest)
						}
						break
					}
				}
				if !hadIntersection {
					seedRanges = append(seedRanges, s)
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
	min := seedRanges[0].start
	for _, s := range seedRanges {
		if min > s.start {
			min = s.start
		}
	}
	fmt.Println(min)
}

func main() {
	PartOne()
	PartTwo()
}
