// с моком!
package main

import "testing"

func Test_person_rudimentaryTranslator(t *testing.T) {
	tests := []struct {
		name string
		p    person
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.rudimentaryTranslator(); got != tt.want {
				t.Errorf("person.rudimentaryTranslator() = %v, want %v", got, tt.want)
			}
		})
	}
}
