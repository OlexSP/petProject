package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "size = 1",
			size: 1,
			want: 1,
		},
		{
			name: "size = 5",
			size: 5,
			want: 5,
		},
		{
			name: "size = 10",
			size: 10,
			want: 10,
		},
		{
			name: "size = 20",
			size: 20,
			want: 20,
		},
		{
			name: "size = 0",
			size: 0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			str2 := NewRandomString(tt.size)
			println(str1)
			// Check that generated strings of the specified length
			assert.Len(t, str1, tt.want)
			assert.Len(t, str2, tt.want)

			// Check that two generated strings are different
			assert.NotEqual(t, str1, str2)

			// Check that all characters are allowed
			for _, char := range str1 {
				if !isAllowedCharacter(char) {
					t.Errorf("Random string contains invalid character: %c", char)
				}
			}
		})
	}
}

func isAllowedCharacter(char rune) bool {
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, allowedChar := range allowedChars {
		if char == allowedChar {
			return true
		}
	}
	return false
}
