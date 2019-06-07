// https://gobyexample.com/goroutines
package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("one")

	go f("two")
	go f("three")

	go func(msg string) {
		fmt.Println(msg)
	}("four")

	fmt.Scanln()
}
