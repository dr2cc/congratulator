package main

import "fmt"

var lang string

func main() {

	t := TranslationService{}

	user := person{
		name:       "Harry",
		Translator: &t,
	}

	lang = "esp"
	fmt.Println("User language is", lang)
	output := user.SayGreeting(lang)
	fmt.Println(output)
}
