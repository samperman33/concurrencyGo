package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()

	fmt.Printf("Going to create %d goruotines.\n", *n)

	var wg sync.WaitGroup

	for i := 0; i < *n; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Wait()
	fmt.Println("quiting...")
}
