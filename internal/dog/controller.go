package dog

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// External interface.
type Controller interface {
	All(http.ResponseWriter, *http.Request)
	ByName(http.ResponseWriter, *http.Request)
}

// External factory.
func NewController(s Service, tpl *template.Template) Controller {
	return &controller{
		service: s,
		tpl:     tpl,
	}
}

// Internal representation.
type controller struct {
	service Service
	tpl     *template.Template
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
	Dogs []*DogResponse
}

func blargh(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	log.Println(err)
}

func (c *controller) render(w http.ResponseWriter, blockName string, blockData any) {
	w.WriteHeader(http.StatusOK)
	if err := c.tpl.ExecuteTemplate(w, blockName, blockData); err != nil {
		log.Printf("c.tpl.ExecuteTemplate(): %v\n", err)
	}
}

// GET /dog/all?order=<name|breed>&direction=<asc|desc>
func (c *controller) All(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		blargh(w, http.StatusInternalServerError, fmt.Errorf("r.ParseForm(): %w", err))

		return
	}

	// construct request
	dr := &AllRequest{
		Order:     r.FormValue("order"),
		Direction: r.FormValue("direction"),
	}

	response, err := c.service.All(dr)
	if err != nil {
		blargh(w, http.StatusInternalServerError, fmt.Errorf("c.service.All(): %w", err))

		return
	}

	c.render(w, "all", response)

}

// Get /dog?name=<name>
func (c *controller) ByName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		blargh(w, http.StatusInternalServerError, fmt.Errorf("r.ParseForm(): %w", err))

		return
	}

	dr := &NameRequest{
		Name: r.FormValue("name"),
	}

	response, err := c.service.GetByName(dr)
	if err != nil {
		if errors.Is(err, ErrNoSuchDog) {
			c.render(w, "no-such-dog", nil)
		} else {
			blargh(w, http.StatusInternalServerError, fmt.Errorf("c.Service.GetByName(%s): %v", dr.Name, err))
		}

		return
	}

	c.render(w, "single", response)
}
