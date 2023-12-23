package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH = iota + 1
	ONE
	TWO
	THREE
	FULL
	FOUR
	FIVE
)

func Part1() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	// map figure to its 'strength'
	labels := make(map[rune]int)
	chars := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for i, l := range chars {
		labels[l] = 20 - i
	}
	countCards := func(hand string) map[rune]int {
		out := make(map[rune]int)
		for _, r := range hand {
			out[r]++
		}
		return out
	}
	handBid := [][]string{}
	handToOrder := make(map[string]int)
	for _, line := range lines {
		split := strings.Split(line, " ")
		hand, bid := split[0], split[1]
		handBid = append(handBid, []string{hand, bid})
		counted := countCards(hand)
		// and ordered by count
		countedPairs := make([][]int, 0)
		for k, v := range counted {
			countedPairs = append(countedPairs, []int{int(k), v})
		}
		sort.Slice(countedPairs, func(i, j int) bool {
			return countedPairs[i][1] > countedPairs[j][1]
		})

		if countedPairs[0][1] == 5 { // five
			handToOrder[hand] = FIVE
		} else if countedPairs[0][1] == 4 { // four
			handToOrder[hand] = FOUR
		} else if countedPairs[0][1] == 3 && countedPairs[1][1] == 2 { // full house
			handToOrder[hand] = FULL
		} else if countedPairs[0][1] == 3 { // three
			handToOrder[hand] = THREE
		} else if countedPairs[0][1] == 2 && countedPairs[1][1] == 2 { // two pair
			handToOrder[hand] = TWO
		} else if countedPairs[0][1] == 2 { // one pair
			handToOrder[hand] = ONE
		} else { // high card
			handToOrder[hand] = HIGH
		}
	}
	sort.Slice(handBid, func(i, j int) bool {
		handI, handJ := handBid[i][0], handBid[j][0]
		if handToOrder[handI] == handToOrder[handJ] {
			for cid := range handI {
				c1, c2 := rune(handI[cid]), rune(handJ[cid])
				if labels[c1] != labels[c2] {
					return labels[c1] < labels[c2]
				}
			}
			return false
		} else {
			return handToOrder[handBid[i][0]] < handToOrder[handBid[j][0]]
		}
	})
	// fmt.Println(handBid)
	sum := 0
	for order, hb := range handBid {
		bid, _ := strconv.Atoi(hb[1])
		sum += (order + 1) * bid
	}
	fmt.Println(sum)
}

func Part2() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	// map figure to its 'strength'
	labels := make(map[rune]int)
	chars := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	for i, l := range chars {
		labels[l] = 20 - i
	}
	countCards := func(hand string) map[rune]int {
		out := make(map[rune]int)
		for _, r := range hand {
			out[r]++
		}
		return out
	}
	handBid := [][]string{}
	handToOrder := make(map[string]int)
	for _, line := range lines {
		split := strings.Split(line, " ")
		hand, bid := split[0], split[1]
		handBid = append(handBid, []string{hand, bid})
		counted := countCards(hand)
		// and ordered by count
		countedPairs := make([][]int, 0)
		for k, v := range counted {
			countedPairs = append(countedPairs, []int{int(k), v})
		}
		sort.Slice(countedPairs, func(i, j int) bool {
			return countedPairs[i][1] > countedPairs[j][1]
		})

		jokerCount := counted['J']

		if countedPairs[0][1] == 5 { // five
			handToOrder[hand] = FIVE
		} else if countedPairs[0][1] == 4 { // four
			if jokerCount > 0 { // either we have JJJJX  or XXXXJ
				handToOrder[hand] = FIVE
			} else { // no jokers
				handToOrder[hand] = FOUR
			}
		} else if countedPairs[0][1] == 3 && countedPairs[1][1] == 2 { // full house
			switch jokerCount {
			case 0:
				handToOrder[hand] = FULL
			case 2, 3: // XXXJJ or JJJXX
				handToOrder[hand] = FIVE
			}
		} else if countedPairs[0][1] == 3 { // three
			switch jokerCount {
			case 0:
				handToOrder[hand] = THREE
			case 1, 3: // we get four if we have XXXJY or JJJXY
				handToOrder[hand] = FOUR
			}
		} else if countedPairs[0][1] == 2 && countedPairs[1][1] == 2 { // two pair (XXYYZ)
			switch jokerCount {
			case 0:
				handToOrder[hand] = TWO
			case 1: // XXYYJ we can turn into full
				handToOrder[hand] = FULL
			case 2: // we get FOUR (XXJJY)
				handToOrder[hand] = FOUR
			}
		} else if countedPairs[0][1] == 2 { // one pair XXABC
			switch jokerCount {
			case 0:
				handToOrder[hand] = ONE
			case 1, 2: // XXJYZ or JJXYZ
				handToOrder[hand] = THREE
			}
		} else { // high card
			switch jokerCount {
			case 0:
				handToOrder[hand] = HIGH
			case 1: // JABCD one pair
				handToOrder[hand] = ONE
			}
		}
	}
	sort.Slice(handBid, func(i, j int) bool {
		handI, handJ := handBid[i][0], handBid[j][0]
		if handToOrder[handI] == handToOrder[handJ] {
			for cid := range handI {
				c1, c2 := rune(handI[cid]), rune(handJ[cid])
				if labels[c1] != labels[c2] {
					return labels[c1] < labels[c2]
				}
			}
			return false
		} else {
			return handToOrder[handI] < handToOrder[handJ]
		}
	})
	sum := 0
	for order, hb := range handBid {
		bid, _ := strconv.Atoi(hb[1])
		sum += (order + 1) * bid
	}
	fmt.Println(sum)
}

func main() {
	Part1()
	Part2()
}
