package robot

import (
	"fmt"
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/volta/config"
	"io/ioutil"
	"net/http"
	"time"
)

const denverURL = `http://www.petharbor.com/results.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=1000&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist=%27ARAP%27,%27AURO%27,%27DNVR%27,%27DDFL%27,%2783615%27,%2780454%27,%2779367%27,%2782294%27,%2777298%27,%2784657%27,%2769972%27,%2784715%27,%2779780%27,%2777803%27,%2776338%27,%2785330%27,%2776065%27,%2778397%27,%2786214%27,%2785252%27,%2774805%27,%2773867%27,%2782242%27,%2781793%27,%2772856%27,%2773086%27,%2782431%27,%2786406%27,%2774867%27,%2783241%27,%2772907%27,%2774328%27,%2786813%27,%2771436%27,%2782755%27,%2782206%27,%2776134%27&atype=&PAGE=1`

var whereTypes = []string{"type_OO", "type_CAT", "type_DOG"}

func GetPets() ([]Pet, error) {
	pets := make([]Pet, 0)

	// Get and parse the results for each animal type
	for _, animal := range whereTypes {
		u := UpdateParameter(denverURL, "where", animal)
		response, err := http.Get(u)
		if err != nil {
			return pets, fmt.Errorf("Error while getting animal %s: %s", err)
		}
		defer response.Body.Close()

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return pets, fmt.Errorf("Error while reading animal %s: %s", err)
		}

		pp, err := ParsePetsHTML(content)
		if pp != nil {
			return pets, fmt.Errorf("Error while parsing animal %s: %s", err)
		}
		pets = append(pets, pp...)
	}
	return pets, nil
}

type handler struct {
	db config.DatabaseConfig
}

func (h handler) updatePetsJob(method func() ([]Pet, error)) error {
	// Get the currently listed pets
	pets, err := method()
	if err != nil {
		return err
	}
	return h.UpdatePets(pets)
}

func (h handler) UpdatePetsJob() error {
	return h.updatePetsJob(GetPets)
}

func (h handler) UpdatePets(pets []Pet) error {
	// Connect to the database
	conn, err := aspect.Connect(h.db.Driver, h.db.Credentials())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get all the current pet ids
	ids := make([]string, len(pets))
	petsByID := make(map[string]Pet)
	for i, pet := range pets {
		ids[i] = pet.ID
		petsByID[pet.ID] = pet
	}

	// Add any pets that don't yet exist and mark any that were removed
	var existing []string
	stmt := aspect.Select(
		Pets.C["id"],
	).OrderBy(
		Pets.C["id"].Asc(),
	)
	if err = conn.QueryAll(stmt, &existing); err != nil {
		return fmt.Errorf("Error while querying existing ids: %s", err)
	}

	newIDs, removedIDs := Complements(existing, ids)

	// Insert new pets into the database
	if len(newIDs) > 0 {
		newPets := make([]PetWithTimestamp, len(newIDs))
		var i int
		for _, id := range newIDs {
			newPets[i] = PWTFromPet(petsByID[id])
			// TODO Database could perform auto-now
			newPets[i].Added = time.Now()
			i += 1
		}
		if _, err = conn.Execute(Pets.Insert(newPets)); err != nil {
			return fmt.Errorf("Error while inserting new pets: %s", err)
		}
	}

	// Update existing pets' removed field (if it hasn't already been set)
	if len(removedIDs) > 0 {
		values := aspect.Values{"removed": time.Now()}
		removeStmt := Pets.Update(values).Where(
			aspect.AllOf(
				Pets.C["id"].In(removedIDs),
				Pets.C["removed"].IsNull(),
			),
		)
		if _, err = conn.Execute(removeStmt); err != nil {
			return fmt.Errorf("Error while updating removed pets: %s", err)
		}
	}
	return nil
}

func NewPetsHandler(db config.DatabaseConfig) (h handler) {
	h.db = db
	return
}
