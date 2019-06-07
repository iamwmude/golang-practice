package main

import (
	"fmt"
	"time"
)

func main() {

	startingTime1 := time.Now().UTC()
	time.Sleep(time.Duration(10) * time.Millisecond)
	endingTime1 := time.Now().UTC()

	fmt.Println(endingTime1.Sub(startingTime1))
	fmt.Println(int64(endingTime1.Sub(startingTime1)))

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	fmt.Println(endingTime.Sub(startingTime) * 65123165)
	fmt.Println(int64(endingTime.Sub(startingTime)))
}
