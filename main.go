package main

import "fmt"

// пользовательский тип данных person
type person struct {
	name     string
	language string
}

func main() {
	//var name string
	//var lang string

	user := person{}

	fmt.Print("What's your name: ")
	fmt.Scan(&user.name)

	fmt.Print("What language do you use? ")
	fmt.Scan(&user.language)

	//fmt.Println("User language is", lang)

	output := user.rudimentaryTranslator()
	fmt.Println(output)

	// // Можно выводить сообщение ниже если язык не поддерживается?!
	// fmt.Printf("Привет, %s! Твой язык %s .\n", name, lang)
	// //
	// // Если переменная имеет тип int то используем параметр %d
	// // var name string
	// // var age int
	// // ...
	// // fmt.Printf("Привет, %s! Тебе %d лет.\n", name, age)
	// //
}
