package service

import "fmt"

// У Жашкевича примеры на mockgen@v1.6.0
// Это старая библиотека, больше не поддерживается google
// Использовать ее не надо (могут возникать коллизии с новой, к примеру не равество типов (по сути одинаковых)
// из старой и новой библиотек)
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

//go:generate mockgen -source=person.go -destination=mocks/mockMockgen.go
type Translator interface {
	RudimentaryTranslator() string
}

// Та, что использовалась по умолчанию- новая версия mockgen
// Ее устанавливать так:
// go install go.uber.org/mock/mockgen@latest
// Создавать файл с моками так:
// mkdir mocks
// mockgen -source person.go > mocks/mock.go
// mockgen -version  вернет версию, сейчас (10.06.2025)- v0.5.2

func Welcome(out Translator) {
	fmt.Println(out.RudimentaryTranslator())
}
