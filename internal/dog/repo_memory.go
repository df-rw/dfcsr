package dog

import (
	"log"
	"slices"
	"sort"
)

type entity struct {
	name  string `json:"name"`
	breed string `json:"breed"`
}

type Repository interface {
	All(*Filters) ([]*Model, error)
}

type dogRepository struct {
	dogs []entity
}

func NewMemoryRepository() Repository {
	return &dogRepository{
		dogs: []entity{
			{"Banjo", "Cocker Spaniel"},
			{"Noah", "Border Collie"},
			{"Sebastian", "Border Collie"},
		},
	}
}

func toModel(entities []entity) []*Model {
	dogs := make([]*Model, len(entities))
	for i, e := range entities {
		dogs[i] = &Model{
			Name:  e.name,
			Breed: e.breed,
		}
	}

	return dogs
}

func (d *dogRepository) sortEntities(filters *Filters) []entity {
	dogs := slices.Clone(d.dogs)

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
			log.Println("name desc")
			sortFn = sortByNameDesc
		}

	case "breed":
		switch filters.Direction {
		case "asc":
			log.Println("breed asc")
			sortFn = sortByBreedAsc

		case "desc":
			log.Println("breed desc")
			sortFn = sortByBreedDesc
		}
	}

	sort.Slice(dogs, sortFn)

	return dogs
}

func (d *dogRepository) All(filters *Filters) ([]*Model, error) {
	dogs := d.sortEntities(filters)

	return toModel(dogs), nil
}
