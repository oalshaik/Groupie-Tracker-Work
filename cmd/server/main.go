package main

import (
	"groupie-trackers/pkg/api"
	"groupie-trackers/pkg/handlers"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	// Handle the main page with preloaded artists
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handlers.RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.")
			return
		}

		// Fetch artists data to display on the main page
		artists, err := api.FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			handlers.RenderErrorPage(w, "Error", "Failed to load artists. Please try again later.")
			return
		}

		// Filter artists if a search query is present
		query := r.URL.Query().Get("search")
		if query != "" {
			filteredArtists := []api.Artist{}
			for _, artist := range artists {
				if containsIgnoreCase(artist.Name, query) {
					filteredArtists = append(filteredArtists, artist)
				}
			}
			artists = filteredArtists
		}

		// Render the main page with the artists data
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, artists) // Pass artists data to the template
	})

	// Handle artist details pages
	mux.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/artists/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			handlers.RenderErrorPage(w, "Invalid Artist ID", "The artist ID provided is not valid.")
			return
		}
		handlers.HandleArtistByID(w, r, id)
	})

	// Handle invalid paths
	mux.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		handlers.RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.")
	})

	// Handle static files (CSS, JS, etc.)
	mux.Handle("/static/", handlers.CheckFileExists(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/")))))

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Utility function to check if a string contains another string, case-insensitively
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
