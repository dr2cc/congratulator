package main

import "fmt"

// Перед go generate нужно установить библиотеку!
// go install github.com/golang/mock/mockgen@v1.6.0
//
// Это старая библиотека, больше не поддерживается google
// У Жашкевича примеры на ней
//
// В этом примере отличие тут:
//
// type MockTranslator struct {
// 	ctrl     *gomock.Controller
// 	recorder *MockTranslatorMockRecorder
// // Эту структуру (isgomock) старая библиотека не создает
// // В данном примере мы ей и не пользуемся
// 	isgomock struct{}
// }

//go:generate mockgen -source=person.go -destination=mocks/mock.go
type Translator interface {
	rudimentaryTranslator(lang string, name string) string
}

// Та, что использовалась по умолчанию- новая версия mockgen
// Ее устанавливать так:
// go install go.uber.org/mock/mockgen@latest
// Создавать файл с моками так:
// mkdir mocks
// mockgen -source person.go > mocks/mock.go
// mockgen -version  вернет версию, сейчас (10.06.2025)- v0.5.2

// // Сервис перевода.
// // В данный момент полностью полагается на метод rudimentaryTranslator
// type translationService struct{}

//функция "простой переводчик" удовлетворяющая интерфейсу Translator
func (p person) rudimentaryTranslator() string {
	switch p.language {
	case "eng":
		return fmt.Sprintf("Glad to see you %s!", p.name)
	case "esp":
		return fmt.Sprintf("Me alegro de verte, %s!", p.name)
	default:
		return fmt.Sprintf("Sorry, I don't speak %s.", p.language)
	}
	//or reach out to an external API here...
}

// //метод "приветствие" типа person
// func (p person) welcome() string {
// 	return rudimentaryTranslator(p.language, p.name)
// 	// switch language {
// 	// case "eng":
// 	// 	return fmt.Sprintf("Hello, my name is %s!", p.name)
// 	// case "esp":
// 	// 	return p.Translator.Translate(language, fmt.Sprintf("Hola, me llamo %s", p.name))
// 	// default:
// 	// 	return fmt.Sprintf("I don't speak %s", language)
// 	// }
// }
