package main

import "fmt"

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

func totalWeight(items []InventoryItem) float64 {
	sum := 0.0
	for _, item := range items {
		sum += item.Weight
	}
	return sum
}

func main() {
	inventory := []InventoryItem{
		{"Алмазный меч", 3.5, false},
		{"Щит", 5.0, false},
		{"Зелье", 0.3, false},
		{"Кольцо силы", 0.1, true},
		{"Зачарованное яблоко", 0.5, false},
	}

	fmt.Println("Инвентарь:")
	for _, item := range inventory {
		fmt.Printf("%s - %.1f кг\n", item.Name, item.Weight)
	}
	fmt.Printf("\nОбщий вес: %.1f кг\n", totalWeight(inventory))
}