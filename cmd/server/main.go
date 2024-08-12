package main

import (
	"html/template"
	"log"
	"net/http"

	"groupie-trackers/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// Handle routes
	r.HandleFunc("/artists", handlers.HandleArtists).Methods("GET")
	r.HandleFunc("/artists/{id}", handlers.HandleArtistByID).Methods("GET")

	// Serve the main HTML page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
