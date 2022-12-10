package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pair struct {
	start1 string
	end1   string
	start2 string
	end2   string
}

func main() {
	pairs := readFile("./puzzle")
	fmt.Println(pairs)
}

func readFile(path string) []Pair {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	var pairs []Pair

	for s.Scan() {
		var pair Pair
		line := s.Text()

		splitByComma := strings.Split(line, ",")
		for i, e := range splitByComma {
			splitByHiphen := strings.Split(e, "-")

			if i == 0 {
				pair.start1 = splitByHiphen[0]
				pair.end1 = splitByHiphen[1]
			} else if i == 1 {
				pair.start2 = splitByHiphen[0]
				pair.end2 = splitByHiphen[1]
			}
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
