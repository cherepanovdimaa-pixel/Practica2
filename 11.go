package main 

import "fmt"

type Product struct {
    Name string
	Category string
	Price float64
}

func filterProducts(product []Product, maxPrice float64, category string) []Product {
	var r []Product
	for _, p := range product {
		if p.Price < maxPrice && p.Category == category {
			r = append(r, p)
		}
	}
	return r
}

func main() {
	products := []Product{
		{"Ноутбук", "Электроника", 50000},
		{"Смартфон", "Электроника", 25000},
		{"Наушники", "Электроника", 3000},
		{"Книга", "Книги", 800},
		{"Ручка", "Канцелярия", 50},
		{"Тетрадь", "Канцелярия", 120},
	}
	filtered := filterProducts(products, 30000, "Электроника")
	fmt.Println("Товары до 30000 руб. из категории Электроника:")
	for _, p := range filtered {
		fmt.Printf("  - %s (%.0f руб.)\n", p.Name, p.Price)
	}
	
	filtered2 := filterProducts(products, 100, "Канцелярия")
	fmt.Println("\nТовары до 100 руб. из категории Канцелярия:")
	for _, p := range filtered2 {
		fmt.Printf("  - %s (%.0f руб.)\n", p.Name, p.Price)
	}
}