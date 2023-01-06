package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie // using slice(db baadme seekh lunga)

func getMovies(w http.ResponseWriter, r *http.Request) {}
func getMovie(w http.ResponseWriter, r *http.Request) {}
func addMovie(w http.ResponseWriter, r *http.Request) {}
func editMovie(w http.ResponseWriter, r *http.Request) {}
func deleteMovie(w http.ResponseWriter, r *http.Request) {}

func main () {
	r := mux.NewRouter()

	// appending to the first get request works
	movies = append(movies, Movie{
		Id: "1",
		Isbn: "216039",
		Title: "Demon Slayer: Infinity train",
		Director: &Director{
			FirstName: "Haruo",
			LastName: "Sotozaki",
		},
	})
	movies = append(movies, Movie{
		Id: "2",
		Isbn: "720149",
		Title: "jujutsu kaisen 0",
		Director: &Director{
			FirstName: "Sunghoo",
			LastName: "Park",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies", editMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
