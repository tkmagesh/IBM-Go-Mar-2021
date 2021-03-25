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

type Products []Product

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

	var products Products = []Product{*p1, *p2, *p3, *p4, *p5}
	products.print()

	//using functions as Methods
	fmt.Println(products.Index(*p3))   // => should print 2
	fmt.Println(products.Include(*p3)) //=> true
	fmt.Println(products.Include(*p6)) //=> false
	fmt.Println("Any Stationary Products ? ", products.Any(func(product Product) bool {
		return product.category == "Stationary"
	})) //=> true
	fmt.Println("Any cost > 50 Products ? ", products.Any(func(product Product) bool {
		return product.cost > 50
	})) //=> false

	fmt.Println("Are all products Stationary? ", products.All(func(product Product) bool {
		return product.category == "Stationary"
	})) //= false

	fmt.Println("Filter Stationary Products ")
	stationaryProducts := products.Filter(func(product Product) bool {
		return product.category == "Stationary"
	})
	//printProducts(stationaryProducts)
	stationaryProducts.print()

}

func (products Products) print() {
	for _, product := range products {
		product.print()
	}
}

func (p Product) print() {
	fmt.Printf("id=%d, name=%s, cost=%v, units=%v, category=%s\n", p.id, p.name, p.cost, p.units, p.category)
}

func (products Products) Index(p Product) int {
	for i, v := range products {
		if v == p {
			return i
		}
	}
	return -1
}

func (products Products) Include(p Product) bool {
	return products.Index(p) != -1
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

func (products Products) Any(predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func (products Products) Filter(predicate func(Product) bool) Products {
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
