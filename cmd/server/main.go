package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-trackers/pkg/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// Handle routes
	mux.HandleFunc("/artists", handlers.HandleArtists)
	mux.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		idStr := r.URL.Path[len("/artists/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}
		// Call the handler with the parsed ID
		handlers.HandleArtistByID(w, r, id)
	})

	// Serve the main HTML page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
