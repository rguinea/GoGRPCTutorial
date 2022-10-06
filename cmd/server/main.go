package main

import (
	"github.com/rguinea/GoGRPCTutorial/internal/db"
	"github.com/rguinea/GoGRPCTutorial/internal/rocket"
	"github.com/rguinea/GoGRPCTutorial/internal/transport/grpc"
	"log"
)

func Run() error {
	// should be responsible for initializing the gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)

	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
