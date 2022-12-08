package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	analyze(readFile("./puzzle"))
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

func analyze(ss []string) {
	for _, s := range ss {
		fmt.Println(len(s))
		if len(s)%2 != 0 {
			fmt.Println("it is not equal")
		}
	}
}
