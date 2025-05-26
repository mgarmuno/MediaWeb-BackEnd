package anime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mgarmuno/MediaWeb-BackEnd/api"
)

// TODO Adapt anime API to new fields in the response from anilist
func SearchAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	api.EnableCors(&w)
	query := r.URL.Query()
	searchString, present := query["searchString"]
	if !present || len(searchString) == 0 {
		http.Error(w, "Query parameter 'searchString' is required in order to preform the anime serach", http.StatusBadRequest)
		return
	}
	err := callAnilistEndpoint(w, searchString[0], "ANIME")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

}

func GetAll(w http.ResponseWriter, r *http.Request) {

}

// TODO probably better in a helper for anilist communications
func callAnilistEndpoint(w http.ResponseWriter, search string, mediaType string) error {
	queryGraphqlFormatted := fmt.Sprintf(queryGraphqlSearch, search, mediaType)
	reqBody := map[string]interface{}{
		"query": queryGraphqlFormatted,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "https://graphql.anilist.co", bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, "Error calling anilist", http.StatusInternalServerError)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error calling anilist", http.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error calling anilist", http.StatusInternalServerError)
		return err
	}

	w.Write(body)

	return nil
}
