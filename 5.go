package main

import ("fmt"
		"strings"
)

func validateUser(name string, age int, email string) error {
    
	if name == ""{
		return fmt.Errorf("Имя не должно быть пустым")
	}
	if len(name) >= 50 {
		return fmt.Errorf("Имя не должно быть длиннее 50 символов")
	}
	if age < 18 || age > 120 {
		return fmt.Errorf("Возраст должен быть от 18 до 120")
	}
	if strings.Index(email, "@") == -1 {
		return fmt.Errorf("Email должен содержать символ @")
	}
	return nil
}

func main() {
    a := validateUser("Дмитрий",18 , "ivan@mail.ru")
	if a != nil {
		fmt.Println("Ошибка:",a)
	} else {
		fmt.Println("Данные введены правильно")
	}
}