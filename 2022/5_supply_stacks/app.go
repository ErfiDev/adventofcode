package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	s.Split(bufio.ScanLines)

	var stacks string

	for s.Scan() {
		t := s.Text()
		if t == "" {
			break
		}

		stacks += t + "\n"
	}

	fmt.Println(stacks)

	_ = strings.NewReader(stacks)

	return nil
}
