package main

import "fmt"


var candidates = []string{"Анна", "Борис", "Виктор"}

func Votes(votes []string) {
	r := make(map[string]int)

	for _, name := range votes {
		r[name]++
	}

	t := len(votes)
	fmt.Println("Результаты голосования:")

	if t == 0 {
		fmt.Println("Голосов не поступало.")
		return
	}

	for _, name := range candidates {
		k := r[name]
		percent := float64(k) / float64(t) * 100
		fmt.Print(name, ": ", k, " голосов (", percent, "%)\n")
	}

	fmt.Println("Всего голосов:", t)
}
func main() {

	votes := []string{
		"Анна", "Борис", "Анна", "Виктор", "Анна",
		"Борис", "Виктор", "Анна", "Борис", "Анна",
	}

	Votes(votes)
}