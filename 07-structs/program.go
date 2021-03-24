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

type PerishableProduct struct {
	Product
	life int
}

func newPerishableProduct(id int, name string, cost float64, units int, category string, life int) *PerishableProduct {
	product := newProduct(id, name, cost, units, category)
	return &PerishableProduct{
		*product,
		life,
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

	mango := newPerishableProduct(9, "Mango", 45, 30, "Fruits", 20)
	fmt.Println(mango)
	fmt.Println(mango.id, mango.name, mango.cost, mango.units, mango.category, mango.life)
}
