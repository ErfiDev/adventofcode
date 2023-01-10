package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

type Stacks [][]string

type Command struct {
	index int
	from  int
	to    int
}

//go:embed puzzle
var puzzle string

func init() {
	puzzle = strings.TrimRight(puzzle, "\n")
	if len(puzzle) == 0 {
		panic("puzzle file are empty!")
	}
}

func main() {
	stacks, commands := readAndSplit(puzzle)

	for _, cmd := range commands {
		for i := 0; i < cmd.index; i++ {
			lastItem := stacks[cmd.from-1][len(stacks[cmd.from-1])-1]
			stacks[cmd.to-1] = append(stacks[cmd.to-1], lastItem)

			stacks[cmd.from-1] = stacks[cmd.from-1][:len(stacks[cmd.from-1])-1]
		}
	}

	for _, s := range stacks {
		fmt.Print(s[len(s)-1])
	}
	fmt.Println()
}

func readAndSplit(pzl string) (Stacks, []Command) {
	splitedPzl := strings.Split(pzl, "\n\n")

	firstPart := splitedPzl[0]
	var data [][]string

	for _, p := range strings.Split(firstPart, "\n") {
		data = append(data, strings.Split(p, ""))
	}

	col, row := len(data[0]), len(data)

	var stacks Stacks

	// extracting stacks
	for i := 0; i <= col-1; i++ {
		if data[row-1][i] != " " {
			stack := []string{}
			for j := row - 2; j >= 0; j-- {
				crateName := data[j][i]
				if crateName != " " {
					stack = append(stack, crateName)
				}
			}

			stacks = append(stacks, stack)
		}
	}

	var commands []Command

	secondPart := splitedPzl[1]
	byLine := strings.Split(secondPart, "\n")

	// extracting commands
	for _, l := range byLine {
		cmd := Command{}

		_, err := fmt.Sscanf(l, "move %d from %d to %d", &cmd.index, &cmd.from, &cmd.to)
		if err != nil {
			log.Fatal(err)
		}

		commands = append(commands, cmd)
	}

	return stacks, commands
}
