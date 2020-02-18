package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tipoRepo string

func welcomeToAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API de tareas (TO-DO)")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Item
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &newTask)

	repo := createRepository()
	repo.CreateItem(newTask)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(createRepository().GetItems())
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var task Item
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &task)
	repo := createRepository()
	repo.UpdateItem(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	repo := createRepository()
	repo.DeleteItem(itemID)
	fmt.Fprintf(w, "La tarea con el ID %v ha sido eliminada correctamente", itemID)

}

func createRepository() TypeRepository {
	if tipoRepo == "mongo" {
		return &MongoRepository{}
	}
	return &MemoryRepository{}
}

func main() {

	tipoRepo = "mongo"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeToAPI)
	router.HandleFunc("/api", createTask).Methods("POST")
	router.HandleFunc("/api", getTasks).Methods("GET")
	router.HandleFunc("/api/{id}", updateTask).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
