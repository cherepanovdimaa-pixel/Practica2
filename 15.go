package main

import "fmt"

type Film struct {
	Name   string
	Year   int
	Score  float64
	Genres []string
}

func findBest(list []Film) Film {
	top := list[0]
	for _, f := range list {
		if f.Score > top.Score {
			top = f
		}
	}
	return top
}

func main() {
	films := []Film{
		{"Форрест Гамп", 1994, 8.8, []string{"драма", "комедия"}},
		{"Матрица", 1999, 8.7, []string{"фантастика", "боевик"}},
		{"Властелин колец", 2001, 8.9, []string{"фэнтези", "приключения"}},
		{"Гладиатор", 2000, 8.5, []string{"боевик", "драма"}},
		{"Зелёная миля", 1999, 9.1, []string{"драма", "фэнтези"}},
	}

	fmt.Println("Лучший фильм по оценке")
	best := findBest(films)
	fmt.Printf("Название: %s\nГод: %d\nОценка: %.1f\n\n", best.Name, best.Year, best.Score)

	fmt.Println("Добавляем жанры к фильму 'Форрест Гамп'")
	films[0].Genres = append(films[0].Genres, "мелодрама")
	fmt.Printf("Теперь жанры: %v\n\n", films[0].Genres)

	fmt.Println("Фильмы в жанре 'драма'")
	for _, f := range films {
		for _, g := range f.Genres {
			if g == "драма" {
				fmt.Printf("- %s (%d), оценка %.1f\n", f.Name, f.Year, f.Score)
				break
			}
		}
	}
}