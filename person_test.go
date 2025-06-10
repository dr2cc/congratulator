package main

import (
	"testing"
)

type fakeTranslationService struct{}

func (t fakeTranslationService) Translate(lang string, text string) string {
	if lang == "esp" {
		return "Hola, me llamo Harry"
	}
	return "Hello, my name is Harry"
}

func TestGreeting(t *testing.T) {
	p := person{
		name:       "Sam",
		Translator: fakeTranslationService{},
	}

	got := p.SayGreeting("esp")
	want := "Hola, me llamo Harry"

	if got != want {
		t.Errorf("Got '%s' but wanted '%s'", got, want)
	}
}
