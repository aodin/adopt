package robot

import (
	"testing"
)

func TestPets(t *testing.T) {
	// Test that the table builds without panicking
	if Pets.Name != "pets" {
		t.Fatalf("Unexpected pets table name: %s", Pets.Name)
	}
}
