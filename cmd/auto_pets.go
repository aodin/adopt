package main

import (
	"flag"
	"github.com/aodin/adopt/robot"
	"github.com/aodin/schedule"
	"github.com/aodin/volta/config"
	"log"
	"os"
	"time"
)

func main() {
	path := flag.String("log", "", "Log File")
	flag.Parse()

	// Append log to a file
	if len(*path) > 0 {
		flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE
		logf, err := os.OpenFile(*path, flags, 0644)
		if err != nil {
			log.Fatalf("Could not set log out: %s", err)
		}
		defer logf.Close()
		log.SetOutput(logf)
	}

	// Create a test database config
	c, err := config.ParseFile("../robot/local_settings.json")
	if err != nil {
		log.Fatalf("Could not load configuration file: %s", err)
	}

	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		log.Fatalf("Could not load timezone")
	}

	// Create a new job handler
	h := robot.NewPetsHandler(c.Database)
	schedule.Daily(
		h.UpdatePetsJob,
		schedule.MustParseClockIn("14:00:00", loc),
	)
	log.Println("Starting update pets job")
	schedule.WaitForJobsToFinish()
}
