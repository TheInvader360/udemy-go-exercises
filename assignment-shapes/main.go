package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{base: 5, height: 5}
	s := square{side: 8}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	//differing implementation
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	//differing implementation
	return s.side * s.side
}

func printArea(s shape) {
	//shared logic
	fmt.Println(s.getArea())
}
