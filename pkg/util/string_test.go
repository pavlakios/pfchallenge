package util

import (
	"strings"
	"testing"
)

// Unit test
func TestRandHexString(t *testing.T) {
	const testLength = 16
	const iterations = 1000
	validHexChars := "0123456789ABCDEF"

	// Helper function to check if a string contains only hex characters
	isValidHex := func(s string) bool {
		for _, c := range s {
			if !strings.ContainsRune(validHexChars, c) {
				return false
			}
		}
		return true
	}

	// Test for correct length and valid hex characters
	for i := 0; i < iterations; i++ {
		result := RandStringHex(testLength)

		// Check length
		if len(result) != testLength {
			t.Errorf("Expected length %d, got %d", testLength, len(result))
		}

		// Check valid characters
		if !isValidHex(result) {
			t.Errorf("Generated string contains invalid characters: %s", result)
		}
	}

	// Test randomness by generating multiple strings and ensuring no duplicates (highly unlikely for large iterations)
	seen := make(map[string]bool)
	for i := 0; i < iterations; i++ {
		result := RandStringHex(testLength)
		if seen[result] {
			t.Errorf("Duplicate string detected: %s", result)
		}
		seen[result] = true
	}
}
