package main

import (
	"flag"
	"github.com/aodin/adopt/robot"
	"github.com/aodin/volta/config"
	"log"
)

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

	// TODO This creates a new database connection for each file
	for _, file := range flag.Args() {
		log.Printf("Parsing %s\n", file)

		if err = h.UpdatePetsFromFile(file); err != nil {
			log.Fatalf("Could not complete update pets job: %s", err)
		}
		// TODO Logging? Output?
	}
}
