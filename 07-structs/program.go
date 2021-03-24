package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float64
	units    int
	category string
}

func newProduct(id int, name string, cost float64, units int, category string) *Product {
	return &Product{
		id:       id,
		name:     name,
		cost:     cost,
		units:    units,
		category: category,
	}
}

func main() {
	var penPtr *Product = newProduct(100, "Pen", 10, 50, "Stationary")
	/*
		var pen Product = *penPtr
		fmt.Println(pen.id, pen.name, pen.cost)
	*/
	fmt.Println(penPtr.id, penPtr.name, penPtr.cost)
	var pencil Product = Product{id: 200, name: "Pencil", cost: 5, units: 200, category: "stationary"}
	fmt.Println(pencil.id, pencil.name, pencil.cost)
}
