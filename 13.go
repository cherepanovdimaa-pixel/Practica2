package main

import "fmt"

func main() {
	expenses := map[string]float64{
		"Еда":         15000,
		"Транспорт":   5000,
		"Развлечения": 3000,
	}

	expenses["Еда"] += 2000

	total := expenses["Еда"] + expenses["Транспорт"] + expenses["Развлечения"]
	fmt.Println("Траты по категориям:")
	fmt.Println("Еда:", expenses["Еда"])
	fmt.Println("Транспорт:", expenses["Транспорт"])
	fmt.Println("Развлечения:", expenses["Развлечения"])
	fmt.Println("Итого:", total)
}
