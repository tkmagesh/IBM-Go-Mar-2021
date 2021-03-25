package main

import (
	"fmt"
	"math"
)

/*
	interface = contract
*/

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type cirlce struct {
	radius float64
}

func (c cirlce) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cirlce) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(shape geometry) {
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}

type myShape struct {
}

func (m myShape) area() float64 {
	return float64(10)
}

func (m myShape) perimeter() float64 {
	return float64(100)
}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)
	c := cirlce{radius: 10}
	measure(c)
	m := myShape{}
	measure(m)
}
