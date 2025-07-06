package main

import (
	"fmt"

	service "github.com/dr2cc/congratulator/internal"
)

// Пользовательский тип данных person .
// Эта структура удовлетворяет (satisfies) интерфейсу
//
//	type translator interface {
//		RudimentaryTranslator() string
//	}
//
// т.к. имеет метод
// func (p person) RudimentaryTranslator() string{}
// не принимающий параметров и возвращающий строку.
type person struct {
	name     string
	language string
}

func main() {

	user := person{}

	fmt.Print("What's your name: ")
	fmt.Scan(&user.name)

	fmt.Print("What language do you use? ")
	fmt.Scan(&user.language)

	// Встраивание (embedding) интерфейса (!?)
	// func welcome(out translator){} принимает translator интерфейс в качестве аргумента,
	// а это означает, что ее можно использовать с любым типом, который удовлетворяет translator интерфейсу.
	// В данном случае мы передаем user переменную (тип person структура),
	// которая удовлетворяет translator интерфейсу, поскольку у нее есть rudimentaryTranslator() метод.
	service.Welcome(user)

}

// метод реализующий RudimentaryTranslator интерфейса Translator
func (p person) RudimentaryTranslator() string {
	switch p.language {
	case "eng":
		return fmt.Sprintf("Glad to see you %s!", p.name)
	case "esp":
		return fmt.Sprintf("Me alegro de verte, %s!", p.name)
	default:
		return fmt.Sprintf("Sorry, I don't speak %s.", p.language)
	}
	//or reach out to an external API here...

	// // Можно выводить сообщение ниже если язык не поддерживается?!
	// fmt.Printf("Привет, %s! Твой язык %s .\n", name, language)
	// //
	// // Если переменная имеет тип int то используем параметр %d
	// // var name string
	// // var age int
	// // ...
	// // fmt.Printf("Привет, %s! Тебе %d лет.\n", name, age)
	// //
}
