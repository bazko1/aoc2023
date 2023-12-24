package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
