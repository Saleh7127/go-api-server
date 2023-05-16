package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Saleh7127/go-api-server/DB"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sort"
)

func AllPlayer(w http.ResponseWriter, r *http.Request) {

	sort.Slice(DB.PlayerDB, func(i, j int) bool {
		return DB.PlayerDB[i].JerseyNumber < DB.PlayerDB[j].JerseyNumber
	})
	fmt.Fprint(w, "All Football player list\n")
	for _, player := range DB.PlayerDB {
		json.NewEncoder(w).Encode(player)
	}
}

func PlayerSearch(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	fmt.Println("Info of the player name", name)
	find := false

	for _, i := range DB.PlayerDB {
		if i.Name == name {
			json.NewEncoder(w).Encode(i)
			find = true
		}
	}
	if find == false {
		fmt.Fprintf(w, "Player is not founded")
	}
	w.WriteHeader(http.StatusOK)
}

func CreateNewEntry(w http.ResponseWriter, r *http.Request) {
	var player DB.Player
	json.NewDecoder(r.Body).Decode(&player)
	DB.PlayerDB = append(DB.PlayerDB, player)
	fmt.Println("A new Player profile is created Successfully")
	w.WriteHeader(http.StatusCreated)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Println("Updating the info ", name)
	for i, p := range DB.PlayerDB {
		if p.Name == name {
			var player DB.Player
			json.NewDecoder(r.Body).Decode(&player)
			DB.PlayerDB[i] = player
			fmt.Println("Info updated successfully")
			return
		}
	}
	fmt.Fprintf(w, "Player in not found in the database")
	w.WriteHeader(http.StatusOK)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Println("Deleting the info ", name)
	for i, p := range DB.PlayerDB {
		if p.Name == name {
			DB.PlayerDB = append(DB.PlayerDB[:i], DB.PlayerDB[i+1:]...)
			fmt.Println("Info Deleted successfully")
			return
		}
	}
	fmt.Fprintf(w, "Player in not found in the database")
	w.WriteHeader(http.StatusAccepted)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Futball Homepage")
}
