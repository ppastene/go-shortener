package keygen

import (
	"strconv"
	"testing"
)

func TestGenerate(t *testing.T) {
	kg := NewKeygen()

	tests := []struct {
		length    uint
		expectErr bool
	}{
		{2, true},
		{3, false},
		{10, false},
		{20, false},
	}

	for _, tt := range tests {
		t.Run("Length: "+strconv.FormatUint(uint64(tt.length), 10), func(t *testing.T) {
			key, err := kg.Generate(tt.length)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
				if len(key) != int(tt.length) {
					t.Errorf("expected key length %d but got %d", tt.length, len(key))
				}
				for _, char := range key {
					if !isValidChar(char) {
						t.Errorf("generated key contains invalid character: %c", char)
					}
				}
			}
		})
	}
}

func isValidChar(c rune) bool {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, validChar := range charset {
		if c == validChar {
			return true
		}
	}
	return false
}
