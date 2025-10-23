package dog

import (
	"errors"
	"fmt"
	"slices"
)

type Model struct {
	Name  string
	Breed string
}

func (m *Model) NameBreed() string {
	return m.Name + m.Breed
}

type Service interface {
	All(*AllRequest) (*AllResponse, error)
	GetByName(*NameRequest) (*DogResponse, error)
}

type dogService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &dogService{
		repo: r,
	}
}

type Filters struct {
	Order     string
	Direction string
}

var (
	validOrders         = []string{"name", "breed"}
	validDirections     = []string{"asc", "desc"}
	ErrInvalidOrder     = errors.New("invalid order")
	ErrInvalidDirection = errors.New("invalid direction")
)

func toFilters(dr *AllRequest) (*Filters, error) {
	if !slices.Contains(validOrders, dr.Order) {
		return nil, fmt.Errorf("%s: %w", dr.Order, ErrInvalidOrder)
	}

	if !slices.Contains(validDirections, dr.Direction) {
		return nil, fmt.Errorf("%s: %w", dr.Direction, ErrInvalidDirection)
	}

	return &Filters{
		Order:     dr.Order,
		Direction: dr.Direction,
	}, nil
}

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
		dogs: dogs,
	}
}

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

func (s *dogService) GetByName(dr *NameRequest) (*DogResponse, error) {
	model, err := s.repo.GetByName(dr.Name)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetByName(%s): %w", dr.Name, err)
	}

	return toDogResponse(model), nil
}
