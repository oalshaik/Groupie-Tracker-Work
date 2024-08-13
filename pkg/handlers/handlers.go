package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-trackers/pkg/api"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Function to render the custom error page
func RenderErrorPage(w http.ResponseWriter, title string, message string) {
	tmpl := template.Must(template.ParseFiles("web/templates/error.html"))
	data := struct {
		Title   string
		Message string
	}{
		Title:   title,
		Message: message,
	}
	tmpl.Execute(w, data)
}

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	log.Printf("Fetching artists data from URL: %s", apiURL)

	artists, err := api.FetchArtists(apiURL)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		RenderErrorPage(w, "Error", "Failed to fetch artists. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func HandleArtistByID(w http.ResponseWriter, r *http.Request, id int) {
	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	log.Printf("Fetching artist data from URL: %s", apiURL)

	artist, err := api.FetchArtistByID(apiURL)
	if err != nil || artist == nil {
		log.Printf("Artist not found or error fetching artist: %v", err)
		RenderErrorPage(w, "Artist Not Found", "The artist you are looking for does not exist.")
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/artist_details.html"))
	tmpl.Execute(w, artist)
}

// New handler for invalid paths
func Handle404(w http.ResponseWriter, r *http.Request) {
	RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.")
}

// Middleware to check file existence
func CheckFileExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "." + r.URL.Path
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			RenderErrorPage(w, "File Not Found", "The file you are looking for is missing or has been renamed.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
