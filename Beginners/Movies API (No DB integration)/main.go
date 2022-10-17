package main

import (
	"encoding/json"
	"math/rand"
	"strconv"

	// Third Part Library
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// Native Library
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// Functions

func getMovies(w http.ResponseWriter, r *http.Request) {

	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Handling Error
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Passing the data to the params variable
	params := mux.Vars(r)

	// Iterating through slice
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Variable
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Randomizing item ID
	movie.ID = strconv.Itoa(rand.Intn(1000000))

	// Append item
	movies = append(movies, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Passing the data to the params variable
	params := mux.Vars(r)

	// Iterate through slice
	for index, item := range movies {
		if item.ID == params["id"] {
			//Delete the movie with the ID that were sent
			movies = append(movies[:index], movies[index+1:]...)
			// Add a new movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			// Randomizing item ID
			movie.ID = strconv.Itoa(rand.Intn(1000000))

			// Append item using postman
			movies = append(movies, movie)
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				return
			}
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	// Set Content Type
	w.Header().Set("Content-Type", "application/json")

	// Passing the data to the params variable
	params := mux.Vars(r)

	// Iterate through slice
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {

	// Mux Router
	r := mux.NewRouter()

	// Populating movies
	movies = append(movies, Movie{})

	// Handlers
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Terminal Return
	fmt.Printf("Server Running at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
