package main

import "fmt"


type Employee struct {
	ID       int     
	Name     string  
	Position string  
	Salary   float64 
}

func c(employees []Employee) (t float64, a float64) {
	if len(employees) == 0 {
		return 0, 0 
	}
	for _, emp := range employees {
		t += emp.Salary
	}
	a = t / float64(len(employees))
	return t, a
}


func PrintEmployee(emp Employee) {
	fmt.Printf("  [%d] %-15s | %-15s | %10.2f руб.\n",
		emp.ID, emp.Name, emp.Position, emp.Salary)
}

func main() {

	employees := []Employee{
		{ID: 1, Name: "Иван Иванов", Position: "Разработчик", Salary: 150000.00},
		{ID: 2, Name: "Пётр Петров", Position: "Тестировщик", Salary: 95000.50},
		{ID: 3, Name: "Анна Сидорова", Position: "Менеджер", Salary: 120000.00},
		{ID: 4, Name: "Ольга Кузнецова", Position: "Дизайнер", Salary: 110000.75},
		{ID: 5, Name: "Дмитрий Смирнов", Position: "DevOps", Salary: 165000.00},
	}

	fmt.Println("Список сотрудников")
	for _, emp := range employees {
		PrintEmployee(emp)
	}

	t, a := c(employees)

	fmt.Println("          Статистика по зарплатам ")
	fmt.Printf("Общий фонд оплаты труда: %10.2f руб.\n", t)
	fmt.Printf("Средняя зарплата:        %10.2f руб.\n", a)
	fmt.Printf("Количество сотрудников:  %d\n", len(employees))

	fmt.Println("       Проверка пустого среза ")
	emptyTotal, emptyAvg := c([]Employee{})
	fmt.Printf("ФОТ: %.2f, Средняя: %.2f\n", emptyTotal, emptyAvg)
}