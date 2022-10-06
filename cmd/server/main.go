package main

import (
	"github.com/rguinea/GoGRPCTutorial/internal/db"
	"github.com/rguinea/GoGRPCTutorial/internal/rocket"
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
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
