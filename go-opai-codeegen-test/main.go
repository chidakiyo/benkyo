package main

import (
	"fmt"
	"github.com/chidakiyo/benkyo/go-opai-codeegen-test/api"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct{}

// Returns all pets
// (GET /pets)
func (s Server) FindPets(w http.ResponseWriter, r *http.Request, params api.FindPetsParams) {

	fmt.Fprintf(w, "GET: /pets")
}

// Creates a new pet
// (POST /pets)
func (s Server) AddPet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "POST: /pets")
}

// Deletes a pet by ID
// (DELETE /pets/{id})

func (s Server) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {

	fmt.Fprintf(w, "DELETE: /pets/{id}")
}

// Returns a pet by ID
// (GET /pets/{id})
func (s Server) FindPetByID(w http.ResponseWriter, r *http.Request, id int64) {

	fmt.Fprintf(w, "GET: /pets/{id}")
}

var _ api.ServerInterface = Server{}

func main() {
	r := chi.NewRouter()
	sv := Server{}
	handle := api.HandlerFromMux(sv, r)
	http.ListenAndServe(":8080", handle)
}
