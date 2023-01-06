package main

import (
	"bufio"
	"log"
	"os"
)

type Stack struct {
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

func main() {
	readFile("./puzzle")
}

func readFile(path string) []Stack {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	// var stacks []string
	// var commands []string

	if s.Scan() {

	}

	return nil
}
