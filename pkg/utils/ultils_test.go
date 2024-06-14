package utils

import (
	"fmt"
	"testing"
)

func TestFormatDomainURL(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"example.com", "https://example.com", nil},
		{"http://example.com", "http://example.com", nil},
		{"https://example.com", "https://example.com", nil},
		{"ftp://example.com", "ftp://example.com", fmt.Errorf("invalid domain format: ftp://example.com")},
		{"not a url", "https://not%20a%20url", fmt.Errorf("parse not a url: first path segment in URL cannot contain colon")},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := FormatDomainURL(tc.input)
			fmt.Println(result)
			if err != nil {
				if tc.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tc.err.Error() {
					t.Errorf("Expected error: %v, but got: %v", tc.err, err)
				}
			} else if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
