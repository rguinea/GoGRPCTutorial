//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/TutorialEdge/go-grpc-services-course/rocket Store
package rocket

import "context"

// Rocket - should contain the definition of our Rocket
type Rocket struct {
	ID     string
	Name   string
	Type   string
	Flight int
}

// Store - Defines the expected interface for the DB implementation
type Store interface {
	GeRocketById(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - the service responsible for updating the rocket identity
type Service struct{}

// New - Returns a new instance of our service
func New() Service {
	return Service{
		Store: store,
	}
}

// GetRocketById - Gets a rocket by id from the store
func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, err) {
	rkt, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, nil
	}
	return rkt, nil
}

// InsertRocket - Inserts a rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, err) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, nil
	}
	return rkt, nil
}

// DeleteRocket - Deletes a rocket from the store
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
