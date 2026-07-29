package main

import (
	"testing"
)

func TestTaskTwo(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		k       int
		expected string
	}{{name: "Стандартная ситуация",
		line:    "aa bb cc aa cc cc cc aa ab ac bb",
		k:       3,
		expected: "cc aa bb"},
		{name: "Пустая строка",
			line:    "",
			k:       3,
			expected: ""},
		{name: "Параметр K больше чем длина слайса line",
			line:    "aa bb cc cc",
			k:       5,
			expected: "cc aa bb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sortByCount(tt.line, tt.k)
			if actual != tt.expected {
				t.Errorf("Получили %q, а ожидали %q", actual, tt.expected)
			}
		})
	}
}
