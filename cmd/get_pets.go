package main

import (
	"github.com/aodin/adopt/robot"
	"github.com/aodin/volta/config"
	"log"
)

// TODO This file only makes sense for bootstrapping a complete set of data
// otherwise subsequent uploads will have their "removed" field set

func main() {
	// Create a test database config
	c, err := config.ParseFile("../robot/local_settings.json")
	if err != nil {
		log.Fatalf("Could not load configuration file: %s", err)
	}

	log.Println("Running update pets job")

	// Create a new job handler
	h := robot.NewPetsHandler(c.Database)
	if err = h.UpdatePetsJob(); err != nil {
		log.Fatalf("Could not complete update pets job: %s", err)
	}
}
