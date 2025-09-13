package main

import (
	"bufio"
	"fmt"
	"os"
)

var PHRASE = "hello"
var MATCH = make(map[string]int)

func first(line []string, in chan<- string) {
	for _, val := range line {
		in <- val
	}
	close(in)
}

func second(in chan<- string, out <-chan string) {
	for word := range out {
		if PHRASE == word {
			MATCH[word]++
			in <- word
		}
	}
	close(in)
}

func printResult(out <-chan string) {
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a txt file!")
		os.Exit(1)
	}

	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("Cannot stat", file)
		os.Exit(1)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open", file)
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}

	fmt.Println(lines)

	A := make(chan string)
	B := make(chan string)

	go first(lines, A)
	go second(B, A)
	printResult(B)

	fmt.Println("Match count:", MATCH[PHRASE])
}
