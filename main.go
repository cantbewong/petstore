package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Pet ...
type Pet struct {
	Sku string `json:"Sku"`
	AnimalType string `json:"AnimalType"`
	Variety string `json:"Variety"`
	Color string `json:"Color"`
}

// Pets ...
var Pets[] Pet

func homePage(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the Pet Store!")
	fmt.Println("Endpoint Hit: Pet Store homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Pets", returnAllPets).Methods("GET")
	myRouter.HandleFunc("/Pet/{Sku}", returnSinglePet).Methods("GET")
	myRouter.HandleFunc("/Pet", createNewPet).Methods("POST")
	myRouter.HandleFunc("/Pet/{Sku}", updatePet).Methods("PUT")
	myRouter.HandleFunc("/Pet/{Sku}", deletePet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnAllPets(w http.ResponseWriter, r * http.Request) {
	fmt.Println("Endpoint Hit: returnAllPets")
	json.NewEncoder(w).Encode(Pets)
}

func returnSinglePet(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	key := vars["Sku"]

	for _,	Pet := range Pets {
		if Pet.Sku == key {
			json.NewEncoder(w).Encode(Pet)
			fmt.Println("Endpoint Hit: Pet found")
		}
	}
}

func createNewPet(w http.ResponseWriter, r * http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Pet Pet
	json.Unmarshal(reqBody, & Pet)
	Pets = append(Pets, Pet)
	json.NewEncoder(w).Encode(Pet)
	fmt.Println("Endpoint Hit: Pet added")
}

func deletePet(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	key := vars["Sku"]

	for index,	Pet := range Pets {
		if Pet.Sku == key {
			Pets = append(Pets[: index], Pets[index + 1: ]...)
			fmt.Println("Endpoint Hit: Pet deleted")
		}
	}
}

func updatePet(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	key := vars["Sku"]

	for index, Nextpet := range Pets {
		if Nextpet.Sku == key {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var newpet Pet
			json.Unmarshal(reqBody, & newpet)
			Pets[index].AnimalType = newpet.AnimalType
			Pets[index].Variety = newpet.Variety
			Pets[index].Color = newpet.Color
			fmt.Println("Endpoint Hit: Pet updated")
			return
		}
	}

	http.Error(w, "Pet not found", http.StatusNotFound)
}

func main() {
	// initialize with a few animals
	Pets = [] Pet {
		Pet {
			Sku: "SKU1",
			AnimalType: "Dog",
			Variety: "poodle",
			Color: "white"},
		Pet {
			Sku: "SKU2",
			AnimalType: "Cat",
			Variety: "Calico",
			Color: "orange"},
		Pet {
			Sku: "SKU3",
			AnimalType: "Bird",
			Variety: "Canary",
			Color: "Yellow"},
	}
	handleRequests()
}
