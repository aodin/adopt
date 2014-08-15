package main

import (
	"flag"
	"github.com/aodin/adopt/robot"
	"github.com/aodin/volta/config"
	"log"
)

// TODO This file only makes sense for bootstrapping a complete set of data
// otherwise subsequent uploads will have their "removed" field set

func main() {
	flag.Parse()

	// Create a test database config
	c, err := config.ParseFile("../robot/local_settings.json")
	if err != nil {
		log.Fatalf("Could not load configuration file: %s", err)
	}

	// Create a new job handler
	h := robot.NewPetsHandler(c.Database)

	log.Printf("Parsing %d files\n", len(flag.Args()))

	// Parse each file
	pets := make([]robot.Pet, 0)
	for _, file := range flag.Args() {
		log.Printf("Parsing %s\n", file)
		pp, err := robot.ParsePetsFile(file)
		if err != nil {
			log.Printf("Error while parsing file %s: %s", file, err)
		}
		pets = append(pets, pp...)
	}
	log.Printf("%d Total Pets\n", len(pets))

	if err = h.UpdatePets(pets); err != nil {
		log.Fatalf("Could not complete update pets job: %s", err)
	}
}
