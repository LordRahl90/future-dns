package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"dns/domains/maths"
	"dns/proto"
	"dns/servers"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup
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

	ms := maths.New(sectorID)
	grpcServer := grpc.NewServer()
	dnsServer := servers.NewGRPCServer(ms)
	proto.RegisterDNSServer(grpcServer, dnsServer)

	wg.Add(1)
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":5500")
		if err != nil {
			panic(err)
		}
		log.Printf("GRPC Server starting on port 5500")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		server := servers.New(sectorID)
		log.Fatal(server.Router.Run(":8080"))
	}()
	wg.Wait()
	// cleanup
	// os.Exit(0)
}
