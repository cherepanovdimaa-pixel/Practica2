package main

import "fmt"


type Order struct {
	ID          int
	Items       []int
	Total       int
	Address     string
	IsCompleted bool
}
var orders = make(map[int]Order)
func a(o Order) {
	orders[o.ID] = o
}

func main() {
	order1 := Order{
		ID:          1,
		Items:       []int{10, 20, 30},
		Total:       150050,
		Address:     "ул. Ленина, 10",
		IsCompleted: false,
	}

	a(order1)

	order2 := Order{
		ID:          2,
		Items:       []int{5, 7},
		Total:       30000,
		Address:     "ул. Мира, 5",
		IsCompleted: true,
	}
	a(order2)

	for id, o := range orders {
		fmt.Println("Заказ", id)
		fmt.Println("Товары:", o.Items)
		fmt.Println("Сумма (коп):", o.Total)
		fmt.Println("Адрес:", o.Address)
		fmt.Println("Выполнен:", o.IsCompleted)
		fmt.Println(" ")
	}
}