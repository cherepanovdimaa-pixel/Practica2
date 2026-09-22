package main

import (
	"fmt"
	"time"
)

type Reading struct {
	Device string
	Temp   float64
	Moist  float64
	When   time.Time
}

func mean(items []Reading) float64 {
	if len(items) == 0 {
		return 0.0
	}

	var sum float64
	for _, r := range items {
		sum += r.Temp
	}

	return sum / float64(len(items))
}

func main() {
	today := time.Now()

	samples := []Reading{
		{Device: "кухня", Temp: 18.3, Moist: 55.0, When: today.Add(-23 * time.Hour)},
		{Device: "спальня", Temp: 25.7, Moist: 38.0, When: today.Add(-19 * time.Hour)},
		{Device: "гостиная", Temp: 20.1, Moist: 60.0, When: today.Add(-15 * time.Hour)},
		{Device: "кухня", Temp: 22.8, Moist: 44.0, When: today.Add(-11 * time.Hour)},
		{Device: "спальня", Temp: 19.5, Moist: 52.0, When: today.Add(-7 * time.Hour)},
		{Device: "гостиная", Temp: 24.2, Moist: 47.0, When: today.Add(-3 * time.Hour)},
	}

	result := mean(samples)

	fmt.Println("Отчёт по датчикам за последние 24 часа:")
	for _, r := range samples {
		fmt.Printf("Комната: %-10s | Тепло: %5.1f°C | Сырость: %5.1f%% | Замер в %s\n",
			r.Device, r.Temp, r.Moist, r.When.Format("15:04"))
	}

	fmt.Printf("\nСреднее тепло по дому: %.2f°C\n", result)
}