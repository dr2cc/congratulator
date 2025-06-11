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
type translator interface {
	RudimentaryTranslator() string
}

// Та, что использовалась по умолчанию- новая версия mockgen
// Ее устанавливать так:
// go install go.uber.org/mock/mockgen@latest
// Создавать файл с моками так:
// mkdir mocks
// mockgen -source person.go > mocks/mock.go
// mockgen -version  вернет версию, сейчас (10.06.2025)- v0.5.2

//метод реализующий rudimentaryTranslator из интерфейса Translator
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

func welcome(out translator) {
	fmt.Println(out.RudimentaryTranslator())
}
