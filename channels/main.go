/*func writeToCh(ch chan int, x int) {
	ch <- x
	close(ch)
}

func main() {
	ch := make(chan int)

	go writeToCh(ch, 123)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	time.Sleep(time.Second)

	_, ok := <-ch
	if ok {
		fmt.Println("ch is opened")
	} else {
		fmt.Println("ch is closed")
	}
}*/

package main

import (
	"fmt"
	"time"
)

func wrToCh(ch chan int, x int) {
	ch <- x
	close(ch)
}

func main() {
	ch := make(chan int)
	var readed int
	go wrToCh(ch, 123)
	time.Sleep(3 * time.Second)
	//fmt.Println(<-ch)
	readed = <-ch
	fmt.Println(readed)

	_, ok := <-ch
	if ok {
		fmt.Println("open")
	} else {
		fmt.Println("close")
	}

}
