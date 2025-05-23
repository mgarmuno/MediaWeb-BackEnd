package main

import (
	"fmt"
	"net/http"

	"github.com/mgarmuno/MediaWeb-BackEnd/api/anime"
	"github.com/mgarmuno/MediaWeb-BackEnd/api/movie"
)

func main() {
	http.HandleFunc("/api/movies/getAll", movie.GetAll)
	http.HandleFunc("/api/anime/getAll", anime.GetAll)
	http.HandleFunc("/api/anime/search", anime.SearchAnime)
	fmt.Println("Serving API...")
	http.ListenAndServe("localhost:8080", nil)
}
