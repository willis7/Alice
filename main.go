package main

import (
	"github.com/willis7/Alice/parser"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/willis7/Alice/clipping"
	"encoding/json"
	"strconv"
)

// Store Clips
var clipStore = make(map[string]clipping.Clipping)

// Variable to generate key for the collection
var id int = 0

//HTTP Post - /api/clippings
func PostClipHandler(w http.ResponseWriter, r *http.Request) {
	var clip clipping.Clipping

	// Decode the incoming Clipping JSON
	err := json.NewDecoder(r.Body).Decode(&clip)
	if err != nil {
		panic(err)
	}
	id++
	k := strconv.Itoa(id)
	clipStore[k] = clip

	j, err := json.Marshal(clip)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//HTTP Get - /api/clippings
func GetClipHandler(w http.ResponseWriter, r *http.Request) {
	var clips []clipping.Clipping
	for _, v := range clipStore {
		clips = append(clips, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(clips)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func main() {
	s := `My Clippings.txt`
	clips := parser.Parse(s)

	// Load the clipStore with initial data obtained from Parse()
	for _, clip := range clips {
		id++
		k := strconv.Itoa(id)
		clipStore[k] = clip
	}

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/clippings", GetClipHandler).Methods("GET")
	r.HandleFunc("/api/clippings", PostClipHandler).Methods("POST")
	//TODO: r.HandleFunc("/api/clippings", GetClipHandler).Methods("PUT")
	//TODO: r.HandleFunc("/api/clippings", GetClipHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
