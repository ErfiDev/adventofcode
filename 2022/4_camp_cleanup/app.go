package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	start1 int
	end1   int
	start2 int
	end2   int
}

func main() {
	pairs := readFile("./puzzle")

	overlap := analyze(pairs)
	fmt.Println(overlap)
}

func analyze(pairs []Pair) int {
	overlap := 0
	for _, p := range pairs {
		if (p.start1 >= p.start2 && p.end1 <= p.end2) || (p.start2 >= p.start1 && p.end2 <= p.end1) {
			overlap++
		}
	}

	return overlap
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
			toInt0, _ := strconv.Atoi(splitByHiphen[0])
			toInt1, _ := strconv.Atoi(splitByHiphen[1])

			if i == 0 {
				pair.start1 = toInt0
				pair.end1 = toInt1
			} else if i == 1 {
				pair.start2 = toInt0
				pair.end2 = toInt1
			}
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
