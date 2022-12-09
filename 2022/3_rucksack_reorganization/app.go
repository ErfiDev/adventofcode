package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	a = 97 // the a character ascii code
	A = 65 // and A character ascii code
)

// the better way to find the common character priorities is finding
// character ascii code and calculate according to a and A character ascii
// then we can find the answer

// for example we got `f` character
// how we can find `f` priority?

// 1: find the `f` character ascii
// 2: if the `f` ascii is higher than `a` ascii code
// we know `f` priority is between 1 and 26
// `f`: 102
// `a`: 97

// `f` - `a` = 5 + 1 = 6
// 6 is `f` priority

func main() {
	allPoints := analyze(readFile("./puzzle"))
	fmt.Println(allPoints)
}

func readFile(path string) []string {
	var stringSlice []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		stringSlice = append(stringSlice, s.Text())
	}

	return stringSlice
}

func analyze(ss []string) int {
	sumOfAll := 0
	for _, s := range ss {
		half := len(s) / 2
		first := s[:half]
		second := s[half:]

		var total = make(map[rune]bool)
		var common string

		for _, f := range first {
			total[f] = true
		}

		for _, s := range second {
			if _, ok := total[s]; ok {
				common = string(s)
			}
		}

		sum := sumOfPriorities(common)
		sumOfAll += sum
	}

	return sumOfAll
}

func getAscii(s string) int {
	return int(s[0])
}

func sumOfPriorities(common string) int {
	ascii := getAscii(common)
	if ascii >= a {
		ascii = (ascii - a) + 1
		return ascii
	} else if ascii < a && ascii >= A {
		ascii = (ascii - A) + 27
		return ascii
	}

	return ascii
}
