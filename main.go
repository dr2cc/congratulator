package main

import "fmt"

// Пользовательский тип данных person .
// Эта структура удовлетворяет (satisfies) интерфейсу
// Translator , т.к. имеет метод rudimentaryTranslator()
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
	// Функция welcome принимает translator интерфейс в качестве аргумента,
	// а это означает, что ее можно использовать с любым типом, который удовлетворяет translator интерфейсу.
	// В данном случае мы передаем user переменную (тип person структура),
	// которая удовлетворяет translator интерфейсу, поскольку у нее есть rudimentaryTranslator() метод.
	welcome(user)

}
