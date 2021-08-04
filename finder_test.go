package finder

import "testing"

func TestFindLongestPrefix(t *testing.T) {
	words := []string{"flop", "florint", "florida"}
	expected := "flo"
	result := FindLongestPrefix(words)
	if result != expected {
		t.Fatalf("Expected: %s, got: %s", expected, result)
	}
}
