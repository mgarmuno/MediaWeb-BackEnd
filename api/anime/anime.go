package anime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Page Page `json:Page`
}

type Page struct {
	Media []Media `json:media`
}

type Media struct {
	ID           int        `json:id`
	Title        Title      `json:title`
	CoverImage   CoverImage `json:coverImage`
	AverageScore int        `json:averageScore`
	Popularity   int        `json:popularity`
	Episode      int        `json:episodes`
	Season       string     `json:season`
	SeasonYear   int        `json:seasonYear`
	IsAdult      bool       `json:isAdult`
}

type CoverImage struct {
	Image string `json:large`
}

type Title struct {
	Romaji  string `json:romaji`
	English string `json:english`
	Native  string `json:native`
}

const queryGraphql = `query {
		Page {
			media(search: "%s", type: %s) {
				id
				title {
					romaji
					english
					native
				}
				coverImage {
					large
				}
				averageScore
				popularity
				episodes
				season
				seasonYear
				isAdult
			}
		}
	}`

func SearchAnime(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
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

func GetAll(rw http.ResponseWriter, rq *http.Request) {

}

// TODO extract to another class for reuse
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

// TODO probably better in a helper for anilist communications
func callAnilistEndpoint(w http.ResponseWriter, search string, mediaType string) error {
	queryGraphqlFormatted := fmt.Sprintf(queryGraphql, search, mediaType)
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

// func callAnilistEndpoint(w http.ResponseWriter, search string, mediaType string) error {

// 	var queryGraphqlReplaced string = fmt.Sprintf(queryGraphql, search, mediaType)
// 	client := graphql.NewClient("https://graphql.anilist.co")
// 	req := graphql.NewRequest(queryGraphqlReplaced)
// 	req.Header.Set("Cache-Control", "no-cache")
// 	ctx := context.Background()

// 	var respData Response
// 	if err := client.Run(ctx, req, &respData); err != nil {
// 		http.Error(w, "Error calling anilist", http.StatusInternalServerError)
// 		return err
// 	}

// 	jsonResp, err := json.Marshal(respData)
// 	if err != nil {
// 		http.Error(w, "Error serializing response from anilist", http.StatusInternalServerError)
// 		return err
// 	}
// 	w.Write(jsonResp)

// 	return nil
// }
