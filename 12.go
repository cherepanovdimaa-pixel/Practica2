package main

import (
	"fmt"
	"strconv"
)

const bin = 2
const dec = 10
const hex = 16

func convert(num string, from int, to int) string {
	n, err := strconv.ParseInt(num, from, 64)
	if err != nil {
		return "ошибка ввода"
	}
	r := strconv.FormatInt(n, to)
	return r
}

func main() {
	var num int64
	fmt.Print("Введите десятичное число: ")
	fmt.Scan(&num)

	fmt.Println("В двоичной:", strconv.FormatInt(num, bin))
	fmt.Println("В hex:", strconv.FormatInt(num, hex))
	fmt.Print("\nВведите число для конвертации: ")
	var s string
	fmt.Scan(&s)
	fmt.Print("Из какой системы (2, 10, 16): ")
	var from int
	fmt.Scan(&from)
	fmt.Print("В какую систему (2, 10, 16): ")
	var to int
	fmt.Scan(&to)
	fmt.Println("Результат:", convert(s, from, to))
}