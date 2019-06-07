package main

import "fmt"

func main() {
	struct2()
}

type s3 struct {
	a int
}

type s2 struct {
	a int
}

type s1 struct {
	a  int
	s2 s2
	s3 *s3
}

func struct2() {
	s11 := s1{
		a:  1,
		s2: s2{a: 2},
		s3: &s3{a: 3},
	}
	passStruct(s11)

	fmt.Println(s11.a, s11.s2.a, s11.s3.a)

	s12 := s1{
		a:  1,
		s2: s2{a: 2},
		s3: &s3{a: 3},
	}
	s12.structFunc()
	fmt.Println(s12.a, s12.s2.a, s12.s3.a)

	s13 := s1{
		a:  1,
		s2: s2{a: 2},
		s3: &s3{a: 3},
	}
	s13.structPointerFunc()
	fmt.Println(s13.a, s13.s2.a, s13.s3.a)
}

func passStruct(s1 s1) {
	s1.a = 11
	s1.s2.a = 22
	s1.s3.a = 33
}

func (s1 *s1) structPointerFunc() {
	s1.a = 11
	s1.s2.a = 22
	s1.s3.a = 33
}

func (s1 s1) structFunc() {
	s1.a = 11
	s1.s2.a = 22
	s1.s3.a = 33
}

func struct1() {
	type person struct {
		name string
		age  int
	}

	fmt.Println(person{name: "XDD"})

	p := struct {
		name string
		age  int
	}{name: "SS"}
	fmt.Println(p)

	g := new(person)
	fmt.Println("g", g)

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	a := person{name: "MooD", age: 27}
	b := person{name: "MooD", age: 27}
	fmt.Println("a", a)
	fmt.Println("a", &a)

	x := a
	fmt.Println("x", x)

	t := &a
	fmt.Println("t", t)
	fmt.Println("t", *t)

	t.age = 12
	fmt.Println("t", t)
	fmt.Println("a", a)
	fmt.Println("x", x)
	fmt.Println("b", b)

	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
