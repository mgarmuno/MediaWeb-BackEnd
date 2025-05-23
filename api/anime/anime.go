package anime

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
)

type coverImage struct {
	large string
}

type title struct {
	romaji  string
	english string
	native  string
}

type media struct {
	id           int
	title        title
	coverImage   coverImage
	averageScore int
	popularity   int
	episodes     int
	season       string
	seasonYear   int
	isAdult      bool
}

type animeResponse struct {
	Page []media
}

func GetAll(rw http.ResponseWriter, rq *http.Request) {

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

	var respData animeResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal("Error callling anilist GraphQL API", err)
	}
}
