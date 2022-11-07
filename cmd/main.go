package main

import (
	"log"
	"os"
	"strconv"

	"dns/servers"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" || env == "development" {
		if err := godotenv.Load(".envs/.env"); err != nil {
			// we need to load the environment values
			// before we can start the applications.
			// Otherwise, we will not read in the correct `SectorID`
			panic(err)
		}
	}

	SectorIDStr := os.Getenv("SECTOR_ID")
	sectorID, err := strconv.ParseFloat(SectorIDStr, 64)
	if err != nil {
		// invalid sectorID is provided, better we panic now than use a
		// wrong sector ID for calculations later on
		panic(err)
	}

	server := servers.New(sectorID)
	log.Fatal(server.Router.Run(":8080"))
}
