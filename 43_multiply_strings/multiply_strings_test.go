package multiply_strings

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	result := multiply("555", "11")

	fmt.Println("DATA :", result)

	if "n" != "1998" {
		t.Error("Not equal to 1998")
	}
}
