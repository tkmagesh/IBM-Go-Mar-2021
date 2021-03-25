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
	/*
		var penPtr *Product = newProduct(100, "Pen", 10, 50, "Stationary")
		fmt.Println(penPtr.id, penPtr.name, penPtr.cost)
		var pencil Product = Product{id: 200, name: "Pencil", cost: 5, units: 200, category: "stationary"}
		fmt.Println(pencil.id, pencil.name, pencil.cost)

		mango := newPerishableProduct(9, "Mango", 45, 30, "Fruits", 20)
		fmt.Println(mango)
		fmt.Println(mango.id, mango.name, mango.cost, mango.units, mango.category, mango.life)
	*/
	p1 := newProduct(6, "Pen", 10, 50, "Stationary")
	p2 := newProduct(5, "Ten", 30, 20, "Grocery")
	p3 := newProduct(4, "Den", 20, 10, "Grocery")
	p4 := newProduct(1, "Len", 50, 30, "Stationary")
	p5 := newProduct(7, "Zen", 40, 20, "Stationary")
	p6 := newProduct(9, "Ken", 40, 20, "Stationary")

	products := []Product{*p1, *p2, *p3, *p4, *p5}
	printProducts(products)
	fmt.Println(Index(products, *p3))   // => should print 2
	fmt.Println(Include(products, *p3)) //=> true
	fmt.Println(Include(products, *p6)) //=> false
	fmt.Println("Any Stationary Products ? ", Any(products, func(product Product) bool {
		return product.category == "Stationary"
	})) //=> true
	fmt.Println("Any cost > 50 Products ? ", Any(products, func(product Product) bool {
		return product.cost > 50
	})) //=> false

	fmt.Println("Are all products Stationary? ", All(products, func(product Product) bool {
		return product.category == "Stationary"
	})) //= false

	fmt.Println("Filter Stationary Products ")
	stationaryProducts := Filter(products, func(product Product) bool {
		return product.category == "Stationary"
	})
	printProducts(stationaryProducts)

}

func printProducts(products []Product) {
	for _, product := range products {
		print(product)
	}
}

func print(p Product) {
	fmt.Printf("id=%d, name=%s, cost=%v, units=%v, category=%s\n", p.id, p.name, p.cost, p.units, p.category)
}

func Index(products []Product, p Product) int {
	for i, v := range products {
		if v == p {
			return i
		}
	}
	return -1
}

func Include(products []Product, p Product) bool {
	return Index(products, p) != -1
}

/* category == "stationary"
category != "stationary"
cost == 30
cost <= 30
cost < 30
cost >= 30
cost > 0
cost > 30 && category == "stationary"

func AnyByCategory(products []Product, category string) bool {
	for _, v := range products {
		if v.category == category {
			return true
		}
	}
	return false
}

func AnyByCost(){

} */

func Any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func Filter(products []Product, predicate func(Product) bool) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/*
Index
Include
Any
All
Filter
*/
