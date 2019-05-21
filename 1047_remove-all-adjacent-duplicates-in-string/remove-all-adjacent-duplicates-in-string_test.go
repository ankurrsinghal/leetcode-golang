package remove_all_adjacent_duplicates_in_string

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	r := removeDuplicates("abbaca")

	if r != "ca" {
		t.Error("Not Equal")
	}
}
