package main

import (
	"fmt"
	"net/http"

	"github.com/mgarmuno/MediaWeb-BackEnd/api/anime"
	"github.com/mgarmuno/MediaWeb-BackEnd/database"
)

func main() {
	database.InitializeDatabase()

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/api/anime/getAll", anime.GetAll)
	http.HandleFunc("/api/anime/search", anime.SearchAnime)
	fmt.Println("Serving API...")
	http.ListenAndServe("localhost:8080", nil)
}
