//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/rguinea/GoGRPCTutorial/internal/rocket Store
package rocket

import (
	"context"
	"log"
)

// Rocket - should contain the definition of our Rocket
type Rocket struct {
	ID     string
	Name   string
	Type   string
	Flight int
}

// Store - Defines the expected interface for the DB implementation
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - the service responsible for updating the rocket identity
type Service struct {
	Store Store
}

// New - Returns a new instance of our service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrieves a rocket from the store by ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	log.Print(ctx)
	rkt, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// AddRocket - Adds a rocket to our store
func (s Service) AddRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	log.Print(ctx)
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket - most likely rapid
// unscheduled disassembly
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	log.Print(ctx)
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
