package main

import "fmt"

const (
	Single = "single"
	Double = "double"
	Suite  = "suite"
)

const (
	Free         = "free"
	Booked       = "booked"
	Maintenance  = "maintenance"
)

type HotelRoom struct {
	RoomType string  
	Status   string  
	Price    float64 
}

var rooms = map[string]HotelRoom{
	"101": {Single, Free, 2500},
	"102": {Double, Free, 4000},
	"201": {Suite, Free, 8000},
	"202": {Single, Maintenance, 2500},
}

func bookRoom(roomNumber string) {
	room, exists := rooms[roomNumber]

	if !exists {
		fmt.Println("Номер", roomNumber, "не найден.")
		return
	}
	if room.Status == Booked {
		fmt.Println("Номер", roomNumber, "уже забронирован.")
		return
	}
	if room.Status == Maintenance {
		fmt.Println("Номер", roomNumber, "на обслуживании, бронирование невозможно.")
		return
	}

	room.Status = Booked
	rooms[roomNumber] = room
	fmt.Println("Номер", roomNumber, "успешно забронирован.")
}

func printRooms() {
	fmt.Println("         Список номеров ")
	for number, room := range rooms {
		fmt.Println("Номер:", number)
		fmt.Println("  Тип:", room.RoomType)
		fmt.Println("  Статус:", room.Status)
		fmt.Println("  Цена:", room.Price, "руб./ночь")
		fmt.Println(" ")
	}
}

func main() {
	printRooms()
	fmt.Println("      Попытки бронирования ")
	bookRoom("101") 
	bookRoom("102") 
	bookRoom("202") 
	bookRoom("101") 
	bookRoom("999") 
	printRooms()
}