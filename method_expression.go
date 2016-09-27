package main

// From this : http://packagemain.blogspot.fr/2016/01/method-expressions.html

import (
	"fmt"
)

func main() {
	o := Order{Id: 1}
	o.AddProduct(Product{Id: 1, Name: "milk", Price: 1.1})

	// --- 1 ---
	// method value
	addProduct := o.AddProduct
	addProduct(Product{Id: 2, Name: "bread", Price: 2.11})

	// Test result
	fmt.Printf("Test 1 : o = %v\n", o)

	// --- 2 ---
	// method expression
	calculateTotalCost := (*Order).CalculateTotalCost

	// the first parameter is the receiver of the method
	cost := calculateTotalCost(&o)
	fmt.Printf("Test 2 : Total Cost : $%f\r\n", cost)

	// --- 3 ----
	// method expression using interface
	calculatePrice := PriceCalculator.CalculateTotalCost

	// the first parameter is always the receiver of the method
	cost = calculatePrice(&o)
	fmt.Printf("Test 3.1 : Total Cost : $%f\r\n", cost)

	o2 := Order{Id: 2}
	o2.AddProduct(Product{Id: 1, Name: "milk", Price: 1.1})

	cost = calculatePrice(&o2)
	fmt.Printf("Test 3.2 : Total Cost : $%f\r\n", cost)
}

type Product struct {
	Id    int32
	Name  string
	Price float32
}

type Order struct {
	Id          int32
	ProductList []Product
}

func (o *Order) AddProduct(p Product) {
	o.ProductList = append(o.ProductList, p)
}

func (o *Order) CalculateTotalCost() (cost float32) {
	for _, p := range o.ProductList {
		cost += p.Price
	}
	return
}

// single method interface

type PriceCalculator interface {
	CalculateTotalCost() (cost float32)
}
