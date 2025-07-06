// с моком!
package main

import (
	"testing"

	mock_main "github.com/dr2cc/congratulator/mocks"

	"go.uber.org/mock/gomock"
)

func Test_person_rudimentaryTranslator(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTranslator := mock_main.NewMocktranslator(ctrl)

	mockTranslator.EXPECT().
		RudimentaryTranslator().
		Return("Glad to see you John!")

	// tests := []struct {
	// 	name string
	// 	p    person
	// 	want string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := tt.p.rudimentaryTranslator(); got != tt.want {
	// 			t.Errorf("person.rudimentaryTranslator() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
