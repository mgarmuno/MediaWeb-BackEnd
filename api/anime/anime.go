package anime

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
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
			media(search: "%s", type: ANIME) {
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
	client := graphql.NewClient("https://graphql.anilist.co")
	queryGraphqlReplaced := fmt.Sprintf(queryGraphql, "Naruto")

	req := graphql.NewRequest(queryGraphqlReplaced)
	req.Header.Set("Cache-Control", "no-cache")
	ctx := context.Background()

	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal("Error callling anilist GraphQL API", err)
	}

	jsonResp, err := json.Marshal(respData)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	w.WriteHeader(http.StatusOK)
}

func GetAll(rw http.ResponseWriter, rq *http.Request) {

}
