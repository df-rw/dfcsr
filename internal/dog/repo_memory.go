package dog

import (
	"errors"
	"slices"
	"sort"
	"strings"
	"time"
)

// External interface.
type Repository interface {
	All(*Filters) ([]*Model, error)
	GetByName(string) (*Model, error)
}

// External factory.
func NewMemoryRepository() Repository {
	times := make([]time.Time, 5)
	times[0], _ = time.Parse(time.DateOnly, "2023-01-01")
	times[1], _ = time.Parse(time.DateOnly, "2024-06-30")
	times[2], _ = time.Parse(time.DateOnly, "2021-04-03")
	times[3], _ = time.Parse(time.DateOnly, "2010-12-31")
	times[4], _ = time.Parse(time.DateOnly, "2015-10-01")

	return &dogRepository{
		dogs: []entity{
			{"Banjo", "Cocker Spaniel", times[0]},
			{"Noah", "Border Collie", times[1]},
			{"Sebastian", "Border Collie", times[2]},
			{"Benny", "Poodle", times[3]},
			{"Growler", "Pit Bull", times[4]},
		},
	}
}

// Internal representation.
type dogRepository struct {
	dogs []entity
}

// Internal
type entity struct {
	name  string    `json:"name"`
	breed string    `json:"breed"`
	dob   time.Time `json:"dob"`
}

// Errors.
var (
	ErrNoSuchDog = errors.New("no such dog")
)

// Convert entity to a model.
func toModel(e entity) *Model {
	return &Model{
		Name:  e.name,
		Breed: e.breed,
		DOB:   e.dob,
	}
}

// Convert slice of entities to slice of models.
func toModels(entities []entity) []*Model {
	dogs := make([]*Model, len(entities))
	for i, e := range entities {
		dogs[i] = toModel(e)
	}

	return dogs
}

// Sort entities.
func sortEntities(dogs []entity, filters *Filters) []entity {
	sortByNameAsc := func(i, j int) bool {
		return dogs[i].name < dogs[j].name
	}
	sortByNameDesc := func(i, j int) bool {
		return dogs[j].name < dogs[i].name
	}
	sortByBreedAsc := func(i, j int) bool {
		return dogs[i].breed < dogs[j].breed
	}
	sortByBreedDesc := func(i, j int) bool {
		return dogs[j].breed < dogs[i].breed
	}
	sortFn := sortByNameAsc

	switch filters.Order {
	case "name":
		switch filters.Direction {
		case "desc":
			sortFn = sortByNameDesc
		}

	case "breed":
		switch filters.Direction {
		case "asc":
			sortFn = sortByBreedAsc

		case "desc":
			sortFn = sortByBreedDesc
		}
	}

	sort.Slice(dogs, sortFn)

	return dogs
}

func (d *dogRepository) All(filters *Filters) ([]*Model, error) {
	entities := slices.Clone(d.dogs)
	dogs := sortEntities(entities, filters)

	return toModels(dogs), nil
}

func (d *dogRepository) GetByName(name string) (*Model, error) {
	for _, dog := range d.dogs {
		if strings.EqualFold(dog.name, name) {
			return toModel(dog), nil
		}
	}

	return nil, ErrNoSuchDog
}
