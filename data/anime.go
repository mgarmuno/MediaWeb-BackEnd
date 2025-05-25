package data

type Anime struct {
	Id            string `clover:"id"`
	AnillistId    int    `clover:"AnilistId"`
	Romaji        string `clover:"romaji"`
	English       string `clover:"english"`
	Native        string `clover:"native"`
	Image         string `clover:"image"`
	Episodes      int    `clover:"episodes"`
	AverageScore  int    `clover:"average_score"`
	Progress      int    `clover:"progress"`
	IsAdult       bool   `clover:"is_adult"`
	Status        string `clover:"status"`
	PersonalScore int    `clover:"personal_score"`
}

type AnimeGenre struct {
	Id      string `clover:"id"`
	AnimeId string `clover:"anime_id"`
	GenreId string `clover:"genre_is"`
}

type Genre struct {
	Id    string `clover:"id"`
	Genre string `clover:"genre"`
}

type AnimeQueryResult struct {
	Anime  Anime
	Genres []Genre
}
