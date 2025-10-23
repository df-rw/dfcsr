package dog

import (
	"errors"
	"fmt"
	"slices"
)

// External interface.
type Service interface {
	All(*AllRequest) (*AllResponse, error)
	GetByName(*NameRequest) (*DogResponse, error)
}

// External factory.
func NewService(r Repository) Service {
	return &dogService{
		repo: r,
	}
}

// Internal representation.
type dogService struct {
	repo Repository
}

// Support structures.
type Filters struct {
	Order     string
	Direction string
}

// Domain model.
type Model struct {
	Name  string
	Breed string
}

// Domain model methods.
func (m *Model) NameBreed() string {
	return m.Name + m.Breed
}

// Errors.
var (
	ErrInvalidOrder     = errors.New("invalid order")
	ErrInvalidDirection = errors.New("invalid direction")
)

// Create a filters request.
func toFilters(dr *AllRequest) (*Filters, error) {
	validOrders := []string{"name", "breed"}
	validDirections := []string{"asc", "desc"}

	filters := &Filters{
		Order:     validOrders[0],
		Direction: validDirections[1],
	}

	if dr.Order != "" {
		if !slices.Contains(validOrders, dr.Order) {
			return nil, fmt.Errorf("%s: %w", dr.Order, ErrInvalidOrder)
		}

		filters.Order = dr.Order
	}

	if dr.Direction != "" {
		if !slices.Contains(validDirections, dr.Direction) {
			return nil, fmt.Errorf("%s: %w", dr.Direction, ErrInvalidDirection)
		}

		filters.Direction = dr.Direction
	}

	return filters, nil
}

// Convert a slice of domain models to a response.
func toAllResponse(m []*Model) *AllResponse {
	dogs := make([]*DogResponse, len(m))

	for i, d := range m {
		dogs[i] = &DogResponse{
			Name:      d.Name,
			Breed:     d.Breed,
			NameBreed: d.NameBreed(),
		}
	}

	return &AllResponse{
		Dogs: dogs,
	}
}

// Convert a single domain model to a response.
func toDogResponse(m *Model) *DogResponse {
	if m != nil {
		return &DogResponse{
			Name:      m.Name,
			Breed:     m.Breed,
			NameBreed: m.NameBreed(),
		}
	}

	return nil
}

// All dogs.
func (s *dogService) All(dr *AllRequest) (*AllResponse, error) {
	filters, err := toFilters(dr)
	if err != nil {
		return nil, fmt.Errorf("toFilters(): %w", err)
	}

	models, err := s.repo.All(filters)
	if err != nil {
		return nil, fmt.Errorf("s.repo.All(): %w", err)
	}

	return toAllResponse(models), nil
}

// Dog by name.
func (s *dogService) GetByName(dr *NameRequest) (*DogResponse, error) {
	model, err := s.repo.GetByName(dr.Name)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetByName(%s): %w", dr.Name, err)
	}

	return toDogResponse(model), nil
}
