package main

import "testing"

func TestIntersecFunc(t *testing.T) {
	tests := []struct {
		name     string
		line1    string
		line2    string
		expected string
	}{
		{name: "Штатная отработка 1", line1: "5 3 4 2 1 6", line2: "6 4 2 4", expected: "4 2 6"},
		{name: "Штатная отработка 2", line1: "1 1 5 3 6", line2: "3 9 9 1 2", expected: "1 3"},
		{name: "Пустое пересечение", line1: "1 2 3", line2: "4 5 6", expected: "Empty intersection"},
		{name: "Невалидный ввод", line1: "1 2 3", line2: "z x c", expected: "Invalid input"},
		{name: "Невалидный ввод", line1: "z x c", line2: "1 2 3", expected: "Invalid input"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := intersection(tt.line1, tt.line2)

			if actual != tt.expected {
				t.Fail()
			}
		})
	}
}
