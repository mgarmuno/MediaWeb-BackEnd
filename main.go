package main

import (
	"fmt"
	"net/http"

	"github.com/mgarmuno/MediaWeb-BackEnd/api/anime"
	"github.com/mgarmuno/MediaWeb-BackEnd/api/movie"
	"github.com/mgarmuno/MediaWeb-BackEnd/data"
)

func main() {
	data.InitializeDatabase()

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/api/movies/getAll", movie.GetAll)
	http.HandleFunc("/api/anime/getAll", anime.GetAll)
	http.HandleFunc("/api/anime/search", anime.SearchAnime)
	fmt.Println("Serving API...")
	http.ListenAndServe("localhost:8080", nil)
}
