package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	A          = Shape{shape: "rock", alias: "A", score: 1}
	B          = Shape{shape: "paper", alias: "B", score: 2}
	C          = Shape{shape: "scissor", alias: "C", score: 3}
	X          = Shape{shape: "rock", alias: "X", score: 1}
	Y          = Shape{shape: "paper", alias: "Y", score: 2}
	Z          = Shape{shape: "scissor", alias: "Z", score: 3}
	win_score  = 6
	draw_score = 3
	lose_score = 0
)

type Shape struct {
	alias string
	score int
	shape string
}

type Round struct {
	op string
	me string
}

func main() {
	rounds := readFile("./puzzle")
	fmt.Println(analyze(rounds))
}

func readFile(path string) []Round {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var rounds []Round

	for s.Scan() {
		a := strings.Split(s.Text(), " ")
		rounds = append(rounds, Round{
			op: a[0],
			me: a[1],
		})
	}

	return rounds
}

func analyze(rounds []Round) int {
	var point int

	for _, r := range rounds {
		if r.op == A.alias && (r.me == X.alias) {
			point += draw_score + X.score
			continue
		} else if r.op == B.alias && (r.me == Y.alias) {
			point += draw_score + Y.score
			continue
		} else if r.op == C.alias && (r.me == Z.alias) {
			point += draw_score + Z.score
			continue
		} else if r.op == A.alias && (r.me == Z.alias) {
			point += lose_score + Z.score
			continue
		} else if r.op == A.alias && (r.me == Y.alias) {
			point += win_score + Y.score
		} else if r.op == B.alias && (r.me == X.alias) {
			point += lose_score + X.score
			continue
		} else if r.op == B.alias && (r.me == Z.alias) {
			point += win_score + Z.score
			continue
		} else if r.op == C.alias && (r.me == X.alias) {
			point += win_score + X.score
			continue
		} else if r.op == C.alias && (r.me == Y.alias) {
			point += lose_score + Y.score
			continue
		}
	}

	return point
}
