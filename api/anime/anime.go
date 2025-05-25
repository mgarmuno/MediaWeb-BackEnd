package anime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	data "github.com/mgarmuno/MediaWeb-BackEnd/data"
)

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

type Response struct {
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
					Native  string `json:"native"`
				} `json:"title"`
				CoverImage struct {
					Large string `json:"large"`
				} `json:"coverImage"`
				AverageScore int      `json:"averageScore"`
				Popularity   int      `json:"popularity"`
				Episodes     int      `json:"episodes"`
				Season       string   `json:"season"`
				SeasonYear   int      `json:"seasonYear"`
				IsAdult      bool     `json:"isAdult"`
				Format       string   `json:"format"`
				Genres       []string `json:"genres"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}

type AnimeApi struct {
	ID            string `json:"id"`
	AnilistId     int
	Romaji        string `json:"romaji"`
	English       string `json:"english"`
	Native        string `json:"native"`
	Image         string `json:"image"`
	Episodes      int    `json:"episodes"`
	AverageScore  int    `json:"averageScore"`
	Progress      int
	IsAdult       bool   `json:"isAdult"`
	Season        string `json:"season"`
	SeasonYear    int    `json:"seasonYear"`
	PersonalScore int
	Status        string
	Genres        []string `json:"genres"`
}

func SearchAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
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

func PostAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	var anime AnimeApi
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&anime); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	animeData := prepareAnimeForData(anime)
	newAnimeId, err := data.InsertAnime(animeData, anime.Genres)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error inserting new anime", http.StatusBadRequest)
	}

	animeData.Id = newAnimeId

	downloadImage(&animeData)
	data.UpdateImageUrl(animeData)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	animeQueryResult := data.GetAllAnime()
	animeResult := prepareAnimeResult(animeQueryResult)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, _ := json.Marshal(animeResult)

	w.Write(body)
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
	var response Response
	errUnmarshal := json.Unmarshal(body, &response)
	if errUnmarshal != nil {
		http.Error(w, "Error unmarshaling the response from anilist", http.StatusInternalServerError)
		return err
	}
	animeToMarshal := prepareAnimeToMarshal(response)
	animeBody, _ := json.Marshal(animeToMarshal)
	w.Write(animeBody)

	return nil
}

func prepareAnimeToMarshal(response Response) []AnimeApi {
	var animeList []AnimeApi

	for _, animeFromResponse := range response.Data.Page.Media {
		var anime AnimeApi
		anime.AnilistId = animeFromResponse.ID
		anime.Romaji = animeFromResponse.Title.Romaji
		anime.English = animeFromResponse.Title.English
		anime.Native = animeFromResponse.Title.Native
		anime.Image = animeFromResponse.CoverImage.Large
		anime.Episodes = animeFromResponse.Episodes
		anime.IsAdult = animeFromResponse.IsAdult
		anime.Season = animeFromResponse.Season
		anime.SeasonYear = animeFromResponse.SeasonYear
		anime.AverageScore = animeFromResponse.AverageScore
		anime.Genres = animeFromResponse.Genres
		animeList = append(animeList, anime)
	}

	return animeList
}

func prepareAnimeResult(animeQueryResult data.AnimeQueryResult) AnimeApi {
	var anime AnimeApi
	anime.ID = animeQueryResult.Anime.Id
	anime.AnilistId = animeQueryResult.Anime.AnillistId
	anime.Romaji = animeQueryResult.Anime.Romaji
	anime.English = animeQueryResult.Anime.English
	anime.Native = animeQueryResult.Anime.Native
	anime.Image = animeQueryResult.Anime.Image
	anime.Episodes = animeQueryResult.Anime.Episodes
	anime.AverageScore = animeQueryResult.Anime.AverageScore
	anime.IsAdult = animeQueryResult.Anime.IsAdult

	anime.Progress = animeQueryResult.Anime.Progress
	anime.Status = animeQueryResult.Anime.Status
	anime.PersonalScore = animeQueryResult.Anime.PersonalScore
	for _, genre := range animeQueryResult.Genres {
		anime.Genres = append(anime.Genres, genre.Genre)

	}

	return anime
}

func prepareAnimeForData(anime AnimeApi) data.Anime {
	var animeData data.Anime
	animeData.AnillistId = anime.AnilistId
	animeData.Romaji = anime.Romaji
	animeData.English = anime.English
	animeData.Native = anime.Native
	animeData.Image = anime.Image
	animeData.Episodes = anime.Episodes
	animeData.AverageScore = anime.AverageScore
	animeData.Progress = anime.Progress
	animeData.IsAdult = anime.IsAdult
	animeData.PersonalScore = anime.PersonalScore
	animeData.Status = anime.Status

	return animeData
}

func downloadImage(animeData *data.Anime) {
	url := animeData.Image

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	var imagePath string = "./img/" + animeData.Id + filepath.Ext(url)

	file, err := os.Create(imagePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println(err)
	}
	animeData.Image = imagePath
}
