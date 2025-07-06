package main

import "testing"

func Test_person_RudimentaryTranslator(t *testing.T) {
	tests := []struct {
		name string
		p    person
		want string
	}{
		{
			name: "eng",
			p: person{
				name:     "John",
				language: "eng",
			},
			want: "Glad to see you John!",
		},
		{
			name: "esp",
			p: person{
				name:     "John",
				language: "esp",
			},
			want: "Me alegro de verte, John!",
		},
		{
			name: "other",
			p: person{
				name:     "John",
				language: "jp",
			},
			want: "Sorry, I don't speak jp.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RudimentaryTranslator(); got != tt.want {
				t.Errorf("person.RudimentaryTranslator() = %v, want %v", got, tt.want)
			}
		})
	}
}
