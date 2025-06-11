package main

import (
	"testing"

	mock_main "congrat/mocks"

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
	// Этот "примерщик" даже работу контроллера не завершил, что обязательно требуют
	defer ctrl.Finish()

	mockTranslator := mock_main.NewMockTranslator(ctrl)

	want := "Hola, me llamo Harry"
	mockTranslator.EXPECT().
		Translate("esp", gomock.Any()).
		Return(want)
		//.AnyTimes()
		// Это вроде как заглушка
		// "которая будет отвечать на вызов любое количество раз и не проваливать тест"

	p := person{
		name:       "Sam",
		Translator: mockTranslator,
	}

	got := p.SayGreeting("esp")
	// // если активировать заглушку выше, то можно добавить строки ниже
	// p.SayGreeting("esp")
	// p.SayGreeting("est")
	// p.SayGreeting("esp")

	if got != want {
		t.Errorf("Got '%s' but wanted '%s'", got, want)
	}
}
