package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Seed struct {
	id    int
	count int
}

func main() {
	seed := readFile("./puzzle")
	analyze(seed)
}

func readFile(path string) []Seed {
	s := make([]Seed, 1)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)

	id := 0

	for r.Scan() {
		if r.Text() == "" {
			id++

			s = append(s, Seed{
				id: id,
			})
			continue
		}

		i, err := strconv.Atoi(r.Text())
		if err != nil {
			log.Fatal(err)
		}

		s[id].count += i
	}

	return s
}

func analyze(seed []Seed) {
	sort.Slice(seed, func(i, j int) bool {
		return seed[i].count > seed[j].count
	})

	fmt.Println(seed[0])
}
