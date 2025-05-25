package data

import (
	"fmt"
	"log"

	clover "github.com/ostafen/clover/v2"
	document "github.com/ostafen/clover/v2/document"
	query "github.com/ostafen/clover/v2/query"
)

func InitializeDatabase() {

	db, _ := clover.Open(DBName)
	fmt.Println("Table created start")
	errAnime := db.CreateCollection(AnimeCol)
	if errAnime != nil {
		log.Println(AnimeCol + " already exists")
	}
	errGenres := db.CreateCollection(GenreCol)
	if errGenres != nil {
		log.Println(GenreCol + " already exists")
	}
	errAnimeGenres := db.CreateCollection(AnimeGenresCol)
	if errAnimeGenres != nil {
		log.Println(AnimeGenresCol + " already exists")
	}

	fmt.Println("Tables creation end")

	defer db.Close()
}

func InsertAnime(anime Anime, genresArray []string) (string, error) {
	db, _ := clover.Open(DBName)
	defer db.Close()

	animeDoc := document.NewDocument()
	err := animeDoc.Unmarshal(anime)
	if err != nil {
		log.Println("Error Anime Unmarshal")
		log.Println(anime)
		return "", err
	}

	newAnimeId, err := db.InsertOne(AnimeCol, animeDoc)
	if err != nil {
		log.Println("Error inserting anime")
		log.Println(anime)
		return "", err
	}
	defer db.Close()

	return newAnimeId, insertGenres(newAnimeId, genresArray)
}

func GetAllAnime() AnimeQueryResult {
	db, _ := clover.Open(DBName)
	var animeResult AnimeQueryResult

	animeDocs, _ := db.FindAll(query.NewQuery(AnimeCol))
	for _, animeDoc := range animeDocs {
		genreDocs, _ := db.FindAll(query.NewQuery(GenreCol))
		var anime Anime
		animeDoc.Unmarshal(&anime)
		animeResult.Anime = anime
		for _, genreDoc := range genreDocs {
			var genre Genre
			genreDoc.Unmarshal(&genre)
			animeResult.Genres = append(animeResult.Genres, genre)
		}
	}

	defer db.Close()

	return animeResult
}

func UpdateImageUrl(anime Anime) {
	db, _ := clover.Open(DBName)

	animeDoc := document.NewDocument()
	animeDoc.Unmarshal(anime)

	db.Update(query.NewQuery(AnimeCol).Where(query.Field("id").Eq(anime.Id)), animeDoc.AsMap())
	// db.UpdateById(AnimeCol, anime.Id, map[string]interface{}{"completed": true})

	defer db.Close()
}

func insertGenres(newAnimeId string, genresArray []string) error {
	db, _ := clover.Open(DBName)
	var animeGenresToInsert []AnimeGenre
	var newGenresToInsert []Genre

	for _, value := range genresArray {
		genre, _ := db.FindFirst(query.NewQuery(GenreCol).Where(query.Field("genre").Eq(value)))
		if genre != nil {
			var animeGenre AnimeGenre
			animeGenre.AnimeId = newAnimeId
			animeGenre.GenreId = genre.Get("genre_id").(string)
			animeGenresToInsert = append(animeGenresToInsert, animeGenre)
		} else {
			var genre Genre
			genre.Genre = value
			newGenresToInsert = append(newGenresToInsert, genre)
		}
	}
	defer db.Close()

	err := insertNewGenres(newGenresToInsert)
	if err != nil {
		return err
	}
	return insertNewAnimeGenre(animeGenresToInsert)
}

func insertNewAnimeGenre(animeGenres []AnimeGenre) error {
	db, _ := clover.Open(DBName)

	for _, animeGenre := range animeGenres {
		animeGenreDoc := document.NewDocument()
		err := animeGenreDoc.Unmarshal(animeGenre)
		if err != nil {
			log.Printf("Error %s unmarshal", AnimeGenresCol)
			log.Println(animeGenre)
			return err
		}
		_, errInsert := db.InsertOne(AnimeGenresCol, animeGenreDoc)
		if errInsert != nil {
			log.Printf("Error inserting new %s", AnimeGenresCol)
			log.Println(animeGenre)
			return err
		}
	}
	defer db.Close()

	return nil
}

func insertNewGenres(genres []Genre) error {
	db, _ := clover.Open(DBName)

	for _, genre := range genres {
		genreDoc := document.NewDocument()
		err := genreDoc.Unmarshal(genre)
		if err != nil {
			log.Printf("Error %s unmarshal", GenreCol)
			log.Println(genre)
			return err
		}
		_, errInsert := db.InsertOne(GenreCol, genreDoc)
		if errInsert != nil {
			log.Printf("Error inserting new %s", GenreCol)
			log.Println(genre)
			return err
		}
	}
	defer db.Close()

	return nil
}
