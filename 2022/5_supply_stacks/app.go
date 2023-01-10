package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

type Stack struct {
	name   string
	crates []Crate
}

type Crate struct {
	name string
}

type Rearrangement struct {
	commands []Command
}

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

	for _, c := range commands {
		
	}
}

func readAndSplit(pzl string) ([]Stack, []Command) {
	splitedPzl := strings.Split(pzl, "\n\n")

	firstPart := splitedPzl[0]
	var data [][]string

	for _, p := range strings.Split(firstPart, "\n") {
		data = append(data, strings.Split(p, ""))
	}

	col, row := len(data[0]), len(data)

	var stacks []Stack

	// extracting stacks
	for i := 0; i <= col-1; i++ {
		if data[row-1][i] != " " {
			stack := Stack{}
			stack.name = data[row-1][i]
			for j := row - 2; j >= 0; j-- {
				name := data[j][i]
				stack.crates = append(stack.crates, Crate{name})
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
