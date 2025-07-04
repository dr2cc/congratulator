package main

import (
	"testing"

	mock_main "github.com/dr2cc/congratulator.git/mocks"

	"go.uber.org/mock/gomock"
)

// type fakeTranslationService struct{}

// // метод типа fakeTranslationService, удовлетворяющий интерфейсу Translator
// func (t fakeTranslationService) Translate(lang string, text string) string {
// 	if lang == "esp" {
// 		return "Hola, me llamo Harry"
// 	}
// 	return "Hello, my name is Harry"
// }

// func TestGreeting(t *testing.T) {
// 	p := person{
// 		name:       "Sam",
// 		Translator: fakeTranslationService{},
// 	}

// 	got := p.SayGreeting("esp")
// 	want := "Hola, me llamo Harry"

// 	if got != want {
// 		t.Errorf("Got '%s' but wanted '%s'", got, want)
// 	}
// }

func TestGreeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	// При использовании Go 1.14+ и GoMock 1.5.0+ вам больше не нужно явно вызывать defer ctrl.Finish()
	// Очистка происходит автоматически через фреймворк тестирования Go
	//defer ctrl.Finish()

	mockTranslator := mock_main.NewMockTranslator(ctrl)

	want := "Hola, me llamo Harry"
	mockTranslator.EXPECT().
		Translate("esp", gomock.Any()).
		Return(want).
		Times(3)
		//.AnyTimes()
		// Это вроде как заглушка
		// "которая будет отвечать на вызов любое количество раз и не проваливать тест"

	p := person{
		name:       "Sam",
		Translator: mockTranslator,
	}

	// // Если несколько методов по порядку
	// gomock.InOrder(
	// 	mock.EXPECT().
	// 	Method(gomock.Any(), "abc").
	// 	Return(123, nil).
	// 	Times(1),
	// 	mock.EXPECT().
	// 	AnotherMethod(gomock.Any(), gomock.Len(3)).
	// 	Return(123, "123"),
	// )

	got := p.SayGreeting("esp")
	// Какой индекс будет в Times(3) , столько раз можем вызвать метод Translate (из person.go)
	p.SayGreeting("esp")
	p.SayGreeting("esp")
	// // если активировать заглушку выше, то можно добавить сколько угодно строк ниже
	// p.SayGreeting("esp")

	if got != want {
		t.Errorf("Got '%s' but wanted '%s'", got, want)
	}
}
