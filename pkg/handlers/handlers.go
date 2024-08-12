package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-trackers/pkg/api"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	log.Printf("Fetching artists data from URL: %s", apiURL)

	artists, err := api.FetchArtists(apiURL)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func HandleArtistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	log.Printf("Fetching artist data from URL: %s", apiURL)

	artist, err := api.FetchArtistByID(apiURL)
	if err != nil {
		log.Printf("Error fetching artist: %v", err)
		http.Error(w, "Failed to fetch artist", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/artist_details.html"))
	tmpl.Execute(w, artist)
}
