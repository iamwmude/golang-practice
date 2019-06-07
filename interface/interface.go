package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, heigh float64
}

type circle struct {
	radius float64
}

func (xdd rect) area() float64 {
	return xdd.width * xdd.heigh
}

func (xdd rect) perim() float64 {
	return xdd.width * xdd.heigh * 10
}

func (mdd circle) area() float64 {
	return mdd.radius
}

func (mdd circle) perim() float64 {
	return mdd.radius * 10
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	var x geometry
	x = rect{width: 5, heigh: 4}
	y := circle{radius: 7}

	measure(x)
	measure(y)

	stringInterface := interface{}("test")
	fmt.Println(stringInterface.(string))

	var intInterface interface{} = 5
	fmt.Println(intInterface.(int))

	// The value of ok is true if the assertion holds.
	// Otherwise it is false and the value of config is the zero value for type int.
	// https://golang.org/ref/spec#Type_assertions
	var assertionTest interface{}
	_, ok := assertionTest.(int)
	fmt.Println("assertionTest", ok)

	boolInterface := interface{}(true)
	fmt.Println(boolInterface.(bool))

	configList := []interface{}{
		"config/cache.ini",
		"config/eslog.ini",
	}

	for _, val := range configList {
		config, ok := val.(string)
		fmt.Println(config, ok)
	}
}
