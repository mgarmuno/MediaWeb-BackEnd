package database

const (
	DBPath = "./mediaweb-database/"
	DBName = "mediaweb.db"

	TablesCreation = `
		CREATE TABLE anime 
		(
		anilist_id     INTEGER     NOT NULL,
		tmdb_id 	   INTEGER  NULL,
		romaji         TEXT NULL    ,
		english        TEXT NULL    ,
		native         TEXT NULL    ,
		image          BLOB    NULL    ,
		episodes       INTEGER     NULL    ,
		average_scrore INTEGER     NULL    ,
		progress       INTEGER     NULL    ,
		personal_score INTEGER     NULL    ,
		PRIMARY KEY (anilist_id)
		);

		CREATE TABLE anime_episodes
		(
		anilist_id    INTEGER      NOT NULL,
		name        TEXT  NOT NULL,
		duration    INTEGER      NULL    ,
		description TEXT  NULL    ,
		date_aired  DATETIME NULL    ,
		PRIMARY KEY (anilist_id, name)
		);

		CREATE TABLE anime_genres
		(
		anilist_id  INTEGER NOT NULL,
		genres_id INTEGER NOT NULL,
		PRIMARY KEY (anilist_id, genres_id)
		);

		CREATE TABLE anime_staff_role
		(
		staff_id INTEGER NOT NULL,
		anilist_id INTEGER NOT NULL,
		role_id  INTEGER NOT NULL,
		PRIMARY KEY (staff_id, anilist_id, role_id)
		);

		CREATE TABLE characters
		(
		anilist_id   INTEGER     NOT NULL,
		char_anilist_id INTEGER     NOT NULL,
		full       TEXT NOT NULL,
		native     TEXT NULL    ,
		role       TEXT NULL    ,
		image      TEXT NULL    ,
		PRIMARY KEY (anilist_id, anilist_id)
		);

		CREATE TABLE genres
		(
		anilist_id INTEGER NOT NULL,
		genre TEXT NOT NULL,
		PRIMARY KEY (anilist_id, genre)
		);

		CREATE TABLE roles
		(
		anilist_id   INTEGER NOT NULL,
		role TEXT NOT NULL,
		PRIMARY KEY (anilist_id, role)
		);

		CREATE TABLE staff
		(
		anilist_id INTEGER     NOT NULL,
		stuff_anilist_id INTEGER    NOT NULL,
		first_name TEXT NULL    ,
		last_name  TEXT NULL    ,
		PRIMARY KEY (anilist_id, stuff_anilist_id)
		);
	`

	AnimeCol       = "anime"
	GenreCol       = "genres"
	AnimeGenresCol = "anime_genres"
	PrefixCover    = "cover-"
	PrefixBanner   = "banner-"
)
