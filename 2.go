package main

import "fmt"

func main() {
    var a float64
	var b float64
	var c float64

	fmt.Println("Калькулятор веса багажа:")
	fmt.Print("Введите вес основного багажа: ")
	fmt.Scan(&a)
	fmt.Print("Введите вес ручной клади: ")
	fmt.Scan(&b)
	fmt.Println("Введите вес доп.ручной клади: ")
	fmt.Scan(&c)

	t := a + b + c
	fmt.Printf("\nОбщий вес багажа:%.2f кг\n", t)
}