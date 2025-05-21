package main

import (
	"fmt"
	"net/http"

	"mediaweb/mediaapi"
)

func main() {
	fmt.Println("Serving ...")
	http.Handle("/movieapi", &mediaapi.MediaAPI{})
	http.ListenAndServe(":8080", nil)
}
