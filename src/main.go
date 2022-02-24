package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Animal struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Legs  int    `json:"legs"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/db_animal?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open", err)
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Animal{})
	handleRequests()

}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/v1/animal", createAnimal).Methods("POST")
	myRouter.HandleFunc("/v1/animal", getAnimal).Methods("GET")
	myRouter.HandleFunc("/v1/animal/{id}", getAnimal).Methods("GET")
	myRouter.HandleFunc("/v1/animal/{id}", updateAnimal).Methods("PUT")
	myRouter.HandleFunc("/v1/animal/{id}", deleteAnimal).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func createAnimal(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var animal Animal
	json.Unmarshal(payloads, &animal)

	db.Create(&animal)

	res := Result{Data: animal, Message: "Success Create Animal", Code: 200}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animaID := vars["id"]

	var animal Animal
	db.First(&animal, animaID)

	res := Result{Data: animal, Message: "Success Get Animal", Code: 200}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func updateAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animaID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var animalUpdate Animal
	json.Unmarshal(payloads, &animalUpdate)

	var animal Animal
	db.First(&animal, animaID)
	db.Model(&animal).Update(animalUpdate)

	res := Result{Data: animal, Message: "Success Update Animal", Code: 200}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animaID := vars["id"]

	var animal Animal

	db.First(&animal, animaID)
	db.Delete(&animal)

	res := Result{Message: "Success Delete Animal", Code: 200}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
