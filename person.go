package main

import "fmt"

type Translator interface {
	Translate(lang string, text string) string
}

type TranslationService struct{}

func (t TranslationService) Translate(lang string, text string) string {
	return text // Reach out to an external API here...
}

type person struct {
	name       string
	Translator Translator
}

func (p person) SayGreeting(language string) string {
	switch language {
	case "eng":
		return fmt.Sprintf("Hello, my name is %s!", p.name)
	case "esp":
		return p.Translator.Translate(language, fmt.Sprintf("Hola, me llamo %s", p.name))
	default:
		return fmt.Sprintf("I don't speak %s", language)
	}
}
