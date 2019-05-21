package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	result := isValid("(({[]}))")

	if !result {
		t.Error("Not Equal")
	}
}

func TestIsValidFalse(t *testing.T) {
	result := isValid("((}{][]))")

	if result {
		t.Error("Equal")
	}
}
