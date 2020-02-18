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

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Isdone bool   `json:"isdone"`
}

var tasks = []todo{}

func welcomeToAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API de tareas (TO-DO)")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask todo
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &newTask)

	if !inSlice(tasks[:], newTask.ID) {
		tasks = append(tasks, newTask)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newTask)
	} else {
		fmt.Fprintf(w, "El ID enviado para generar ya existe en la lista de tareas")
	}

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var task todo
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &task)

	if inSlice(tasks[:], itemID) {
		for i, v := range tasks {
			if v.ID == itemID {
				v.Isdone = task.Isdone
				v.Title = task.Title
				tasks = append(tasks[:i], v)
				json.NewEncoder(w).Encode(tasks)
			}
		}
	} else {
		fmt.Fprintf(w, "El ID %d enviado para generar NO existe en la lista de tareas", task.ID)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if inSlice(tasks[:], itemID) {
		for i, v := range tasks {
			if v.ID == itemID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				fmt.Fprintf(w, "La tarea con el ID %v ha sido eliminada correctamente", itemID)
			}
		}
	} else {
		fmt.Fprintf(w, "El ID enviado para generar NO existe en la lista de tareas")
	}
}

func inSlice(arr []todo, val int) bool {
	for _, v := range arr {
		if v.ID == val {
			return true
		}
	}
	return false
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeToAPI)
	router.HandleFunc("/api", createTask).Methods("POST")
	router.HandleFunc("/api", getTasks).Methods("GET")
	router.HandleFunc("/api/{id}", updateTask).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
