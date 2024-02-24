package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"math/rand"
	"net/http"

	// "strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// GET ALL THE MOVIES FUNCTION
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// DELETE A MOVIE
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// code here
	w.Header().Set("Cotent-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// GET A MOVIE FUNCTION
func getMovie(w http.ResponseWriter, r *http.Request) {
	// code here
	w.Header().Set("Conetent-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

// CREATE A MOVIE
func createMovie(w http.ResponseWriter, r *http.Request) {
	// code here
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// UPDATE A MOVIE
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// code here
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1234567", Title: "Movie One Title", Director: &Director{Firstname: "Joshua", Lastname: "Biong"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7654321", Title: "Movie Two Title", Director: &Director{Firstname: "Lywy", Lastname: "Biong"}})

	//========ROUTES=================
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("Get")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server start in port 8000\n")
	log.Fatal(http.ListenAndServe("locahost:8000", r))

}
