package main

import (
	"fmt"
	"strings"
)

type TextStats struct {
	Symbol  int
	Words   int
	Predloz int
}

func textStats(text string) TextStats {
	symbolCount := len([]rune(text))
	words := strings.Fields(text)
	wordCount := len(words)

	predlozCount := 0
	for _, ch := range text {
		if ch == '.' || ch == '!' || ch == '?' {
			predlozCount++
		}
	}

	return TextStats{
		Symbol:  symbolCount,
		Words:   wordCount,
		Predloz: predlozCount,
	}
}

func main() {
	text := "Я пишу эту программу. Надеюсь будет работать! Скоро узнаем?"
	stats := textStats(text)
	fmt.Println("Статистика текста:")
	fmt.Printf("Символов: %d\n", stats.Symbol)
	fmt.Printf("Слов: %d\n", stats.Words)
	fmt.Printf("Предложений: %d\n", stats.Predloz)
}