package robot

import (
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/volta/config"
	"testing"
)

func examplePets() ([]Pet, error) {
	pets := []Pet{
		Pet{ID: "1", Name: "Wolvy"},
		Pet{ID: "2", Name: "Sparty"},
	}
	return pets, nil
}

func lessPets() ([]Pet, error) {
	pets := []Pet{
		Pet{ID: "2", Name: "Sparty"},
	}
	return pets, nil
}

func TestUpdatePetsJob(t *testing.T) {
	// Create a test database config
	c, err := config.ParseFile("./test_settings.json")
	if err != nil {
		t.Fatalf("Could not load configuration file: %s", err)
	}

	// Create a new job handler
	h := NewPetsHandler(c.Database)

	// Use the example pets
	if err = h.updatePetsJob(examplePets); err != nil {
		t.Fatalf("Could not complete update pets job: %s", err)
	}

	// Connect to the database
	conn, err := aspect.Connect(c.Database.Driver, c.Database.Credentials())
	if err != nil {
		t.Fatalf("Could not connect to the database: %s", err)
	}
	defer conn.Close()

	// Re-select the pets
	var pets []PetWithTimestamp
	if err = conn.QueryAll(Pets.Select(), &pets); err != nil {
		t.Fatalf("Could not query pets: %s", err)
	}
	if len(pets) != 2 {
		t.Fatalf("Unexpected number of pets: %d", len(pets))
	}

	// Delete one of the pets
	if _, err = conn.Execute(Pets.Delete().Where(Pets.C["id"].Equals(1))); err != nil {
		t.Fatalf("Could not delete pet with id 1: %s", err)
	}

	// Repeat the original upload
	if err = h.updatePetsJob(examplePets); err != nil {
		t.Fatalf("Could not repeat update pets job: %s", err)
	}

	// Update with less pets
	if err = h.updatePetsJob(lessPets); err != nil {
		t.Fatalf("Could not perform less pets job: %s", err)
	}

	// Re-select the pets
	var petsAgain []PetWithTimestamp
	if err = conn.QueryAll(Pets.Select().OrderBy(Pets.C["id"]), &petsAgain); err != nil {
		t.Fatalf("Could not query pets again: %s", err)
	}
	if len(petsAgain) != 2 {
		t.Fatalf("Unexpected number of pets again: %d", len(petsAgain))
	}

	// Wolvy should have a removed time, but not sparty
	if petsAgain[0].Removed == nil {
		t.Fatalf("No removal time for Wolvy: %s", petsAgain[0].Removed)
	}
	if petsAgain[1].Removed != nil {
		t.Fatalf("Unexpected removal for Sparty: %s", petsAgain[1].Removed)
	}

	// Clean up the database
	// TODO clean-up should always be performed - either use defer or
	// a transaction
	if _, err = conn.Execute(Pets.Delete()); err != nil {
		t.Fatalf("Could not delete pets: %s", err)
	}
}
