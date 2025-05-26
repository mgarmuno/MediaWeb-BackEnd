package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mgarmuno/MediaWeb-BackEnd/api/anime"
)

func InitializeDatabase() {
	checkDatabaseExists()
}

func InsertAnime(anime anime.InfoResponse) {

}

func checkDatabaseExists() {
	if _, err := os.Stat(DBPath + DBName); os.IsNotExist(err) {
		fmt.Println("The database does not exists, creating...")
		createDatabaseStructure()
	}
}

func createDatabaseStructure() {
	db, err := sql.Open("sqlite3", DBPath+DBName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, errtc := db.Exec(TablesCreation)
	if errtc != nil {
		log.Panic("ERROR CREATING TABLES", errtc)
	}

	log.Println("Database and table created")
}
