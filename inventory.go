
package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Stock    int
	Price    float64
}


func DeductStock(product *Product, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if quantity > product.Stock {
		return errors.New("insufficient stock")
	}

	product.Stock -= quantity

	return nil
}



func Restock(product *Product, quantity int) error {
	if quantity <=0 {
		return errors.New("Quantity must be greater than 0 to restock")
	}
	product.Stock += quantity

	return nil
}

// TODO 3: Implement ApplyDiscount.
//
// Requirements:
// - Accept a pointer to Product.
// - Accept a discount percentage.
// - Discount must be between 0 and 100.
// - Modify the original Product.Price.
//
// Example:
// Price = 1000, discount = 10
// New Price = 900

func ApplyDiscount(product *Product, percentage float64) error {
	if percentage < 0 || percentage > 100 {
		return errors.New("discount percentage must be between 0 and 100")
	}

	discountAmount := product.Price * (percentage / 100)
	product.Price -= discountAmount


	return nil
}

func inventory() {
	product := Product{
		ID:    101,
		Name:  "Mechanical Keyboard",
		Stock: 50,
		Price: 2500,
	}

	fmt.Println("Initial Product:", product)

	// Test 1: Purchase 5 keyboards.
	err := DeductStock(&product, 5)
	fmt.Println("Purchase 5:", err)
	fmt.Println("After purchase:", product)

	// Test 2: Try purchasing 100 keyboards.
	err = DeductStock(&product, 100)
	fmt.Println("Purchase 100:", err)
	fmt.Println("After failed purchase:", product)

	// Test 3: Restock 30 keyboards.
	err = Restock(&product, 30)
	fmt.Println("Restock 30:", err)
	fmt.Println("After restocking:", product)

	// Test 4: Apply 10% discount.
	err = ApplyDiscount(&product, 10)
	fmt.Println("Apply discount:", err)
	fmt.Println("After discount:", product)
}
