package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const LENGTH = 5

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func inChannel(min int, max int, out chan<- int) {
	for i := 0; i < LENGTH; i++ {
		out <- random(min, max)

	}
	close(out)
}

func sumChannel(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printSum(in <-chan int) {
	var sum int
	for x2 := range in {
		fmt.Printf("%d/%f ", x2, math.Sqrt(float64(x2)))
		sum += x2
	}
	fmt.Println()
	fmt.Println(sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid input")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Println("error: min > min")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan int)
	ch2 := make(chan int)

	go inChannel(n1, n2, ch1)
	go sumChannel(ch1, ch2)
	printSum(ch2)
}
