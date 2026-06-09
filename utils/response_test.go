package utils

import "testing"

func TestNewError(t *testing.T) {
	got := NewError("not found", 404)

	if got.Error != "not found" {
		t.Errorf("Error = %q, want %q", got.Error, "not found")
	}
	if got.Status != 404 {
		t.Errorf("Status = %d, want %d", got.Status, 404)
	}
}