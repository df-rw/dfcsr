package dog

import (
	"log"
	"net/http"
)

type Controller interface {
	All(http.ResponseWriter, *http.Request)
}

type controller struct {
	service Service
}

func NewController(s Service) Controller {
	return &controller{
		service: s,
	}
}

type AllRequest struct {
	Order     string
	Direction string
}

type DogResponse struct {
	Name      string
	Breed     string
	NameBreed string
}

type AllResponse struct {
	dogs []*DogResponse
}

// GET /dog/all?order=name&direction=asc
func (c *controller) All(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("r.ParseForm()", err)

		return
	}

	// construct DTO
	dr := &AllRequest{
		Order:     r.FormValue("order"),
		Direction: r.FormValue("direction"),
	}

	response, err := c.service.All(dr)
	if err != nil {
		log.Println("c.service.All()", err)

		return
	}

	w.WriteHeader(http.StatusOK)

	for _, d := range response.dogs {
		log.Printf("%s: %s: %s\n", d.Name, d.Breed, d.NameBreed)
	}
	log.Println()
}
