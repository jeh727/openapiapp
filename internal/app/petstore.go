package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jeh727/openapiapp/internal/app/appapi"
)

type PetStore struct{}

// FindPets Returns all pets (GET /pets).
func (*PetStore) FindPets(w http.ResponseWriter, _ *http.Request, params appapi.FindPetsParams) {
	fmt.Fprintf(w, "FindPets called with params: %+v\n", params)
}

// AddPet Creates a new pet (POST /pets).
func (*PetStore) AddPet(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "AddPet called, %+v\n", string(body))
}

// DeletePet Deletes a pet by ID (DELETE /pets/{id}).
func (*PetStore) DeletePet(w http.ResponseWriter, _ *http.Request, id int64) {
	fmt.Fprintf(w, "DeletePet called with ID: %d\n", id)
}

// FindPetByID returns a pet by ID (GET /pets/{id}).
func (*PetStore) FindPetByID(w http.ResponseWriter, _ *http.Request, id int64) {
	fmt.Fprintf(w, "FindPetByID called with ID: %d\n", id)
}
