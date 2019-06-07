package main

import "fmt"

var test = 1

func init() {
	fmt.Println("2", test)
}

func init() {
	fmt.Println("1", test)
	defer fmt.Println("5", test)
	test = 2
}

func main() {
	test = 3
	defer fmt.Println("3", test)
	fmt.Println("MooD", test)
	test = 4
	defer fmt.Println("4", test)
	test = 5
}
