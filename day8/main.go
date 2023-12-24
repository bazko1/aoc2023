package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Part1() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	instructions := lines[0]
	tree := make(map[string][]string)
	for _, line := range lines[2 : len(lines)-1] {
		nc := strings.Split(line, " = ")
		node := nc[0]
		childs := strings.Split(nc[1][1:len(nc[1])-1], ", ")
		tree[node] = childs
	}
	current := "AAA"
	steps := 0
pathloop:
	for {
		for _, instr := range instructions {
			if current == "ZZZ" {
				break pathloop
			}
			childs := tree[current]
			switch instr {
			case 'L':
				current = childs[0]
			case 'R':
				current = childs[1]
			}
			steps += 1
		}
	}
	fmt.Println(steps)
}

func Part2() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	instructions := lines[0]
	tree := make(map[string][]string)
	begins := []string{}
	for _, line := range lines[2 : len(lines)-1] {
		nc := strings.Split(line, " = ")
		node := nc[0]
		childs := strings.Split(nc[1][1:len(nc[1])-1], ", ")
		tree[node] = childs
		if node[len(node)-1] == 'A' {
			begins = append(begins, node)
		}
	}

	findEndAndCycle := func(begin string) int {
		current := begin
		steps := 0
		for {
			for _, instr := range instructions {
				if current[len(current)-1] == 'Z' {
					return steps
				}
				childs := tree[current]
				switch instr {
				case 'L':
					current = childs[0]
				case 'R':
					current = childs[1]
				}
				steps += 1
			}
		}
	}

	cycles := []int{}

	for _, b := range begins {
		it := findEndAndCycle(b)
		cycles = append(cycles, it)
	}
	gcd := cycles[0]
	for _, c := range cycles[1:] {
		gcd = Gcd(gcd, c)
	}

	mult := uint64(cycles[0])
	for _, c := range cycles[1:] {
		mult *= uint64(c) / uint64(gcd)
	}
	fmt.Println(mult)
}
func main() {
	Part1()
	Part2()
}
