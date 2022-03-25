package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

type ReturnDefault struct {
	Message string `json:"message,omitempty"`
}

var people = []Person{
	{ID: "1", Firstname: "Kay", Lastname: "Sampa", Address: &Address{Country: "Brasil", State: "SP"}},
	{ID: "2", Firstname: "Ronaldo", Lastname: "Fenomeno", Address: &Address{Country: "Brasil", State: "SP"}},
	{ID: "3", Firstname: "Cristiano", Lastname: "Ronaldo", Address: &Address{Country: "Portugal", State: "SP"}},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(ReturnDefault{Message: "Deu Ruim"})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		log.Fatal("Error to create person")
	}

	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}
