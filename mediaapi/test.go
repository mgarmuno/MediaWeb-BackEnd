package mediaapi

import (
	"fmt"
	"log"
	"net/http"
)

type MediaAPI struct{}

func (test *MediaAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to %v\n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "ok")
}
