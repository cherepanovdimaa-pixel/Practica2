package main

import "fmt"

func Tags(postsTags [][]string) map[string]bool {

	uniqueTags := make(map[string]bool)
	for _, tags := range postsTags {
		for _, tag := range tags {

			uniqueTags[tag] = true
		}
	}
	return uniqueTags
}

func main() {
	postsTags := [][]string{
		{"go", "backend", "go"},
		{"git", "go", "tools"},
		{"python", "backend", "c#"},
		{"go", "tools", "git"},
	}

	uniqueTags := Tags(postsTags)

	fmt.Println("Уникальные теги:")
	for tag := range uniqueTags {
		fmt.Println(" -", tag)
	}
	fmt.Printf("\nВсего уникальных тегов: %d\n", len(uniqueTags))
}