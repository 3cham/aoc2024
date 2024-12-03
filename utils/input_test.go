package utils

import "testing"

func TestGetInput(t *testing.T) {
	input := GetInput("1", false)
	if input[0] != "174" {
		t.Errorf("GetInput() = %v, want %v", input, "test")
	}
}
