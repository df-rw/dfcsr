package dog

import (
	"log"
	"net/http"
)

// External interface.
type Controller interface {
	All(http.ResponseWriter, *http.Request)
	ByName(http.ResponseWriter, *http.Request)
}

// External factory.
func NewController(s Service) Controller {
	return &controller{
		service: s,
	}
}

// Internal representation.
type controller struct {
	service Service
}

// Request structures.
type AllRequest struct {
	Order     string
	Direction string
}

type NameRequest struct {
	Name string
}

// Response structures.
type DogResponse struct {
	Name      string
	Breed     string
	NameBreed string
}

type AllResponse struct {
	dogs []*DogResponse
}

// GET /dog/all?order=<name|breed>&direction=<asc|desc>
func (c *controller) All(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("r.ParseForm()", err)

		return
	}

	// construct request
	dr := &AllRequest{
		Order:     r.FormValue("order"),
		Direction: r.FormValue("direction"),
	}

	response, err := c.service.All(dr)
	if err != nil {
		log.Printf("c.service.All(): %v", err)

		return
	}

	w.WriteHeader(http.StatusOK)

	for _, d := range response.dogs {
		log.Printf("%s: %s: %s\n", d.Name, d.Breed, d.NameBreed)
	}
	log.Println()
}

// Get /dog?name=<name>
func (c *controller) ByName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("r.ParseForm()", err)

		return
	}

	dr := &NameRequest{
		Name: r.FormValue("name"),
	}

	response, err := c.service.GetByName(dr)
	if err != nil {
		log.Printf("c.Service.GetByName(%s): %v", dr.Name, err)

		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println(response)
}
