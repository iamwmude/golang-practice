package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := make([]int, 0)

	chan1 := make(chan int)
	go func() {
		for {
			arr = append(arr, <-chan1)
		}
	}()

	// go func() {
	// 	time.Sleep(time.Millisecond)
	// 	for i := 0; i < 5; i++ {
	// 		chan1 <- i
	// 	}
	// }()

	j := 0

	addJ := func() {
		j++
	}
	mutex := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			mutex.Lock()
			addJ()
			mutex.Unlock()

			chan1 <- j
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(arr)
}
