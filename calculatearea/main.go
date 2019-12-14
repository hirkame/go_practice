package main

import "fmt"

type shape interface {
	getArea() float64
}

type squere struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func main() {
	t := triangle{base: 11, height: 12}
	s := squere{sideLength: 10}

	printArea(t)
	printArea(s)
}

func (sq squere) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
