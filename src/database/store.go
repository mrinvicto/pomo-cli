package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrinvicto/pomo-cli/src/config"
	"github.com/mrinvicto/pomo-cli/src/models"
)

type Database interface {
	InitDB()
	CreateTables()
	CreateSession(s *models.Session)
	UpdateSession(s *models.Session)
}

type database struct {
	config config.PomoCliConfig
	db     *sql.DB
}

var databaseInstance Database

func GetDatabase() Database {
	return databaseInstance
}

func InitDatabase(c config.PomoCliConfig) Database {
	databaseInstance = &database{
		config: c,
	}

	return databaseInstance
}

func (d *database) InitDB() {
	makeDir(d.config.GetDataStoreFileFolderPath())
	db := establishConnection(d.config.GetDataStoreFileCompletePath())
	fmt.Println("DB Init done")
	d.db = db
}

func (d *database) CreateSession(s *models.Session) {
	result, err := d.db.Exec(
		"INSERT INTO session (title, tags, duration, status) VALUES (?, ?, ?, ?)",
		s.Title, string(s.GetTagsJSON()), s.Duration, s.Status,
	)

	if err != nil {
		log.Fatal(err)
	}

	sessionID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	s.ID = int(sessionID)
}

func (d *database) UpdateSession(s *models.Session) {

}

func (d *database) CreateTables() {
	createTableSession(d.db)
}

func makeDir(dbStoreDirectoryPath string) {
	os.MkdirAll(dbStoreDirectoryPath, os.ModePerm) // Ensure directory exists
}

func establishConnection(sqlLiteFilePath string) *sql.DB {
	db, err := sql.Open("sqlite3", sqlLiteFilePath)

	if err != nil {
		log.Fatal("An error occurred while creating DB connection", err)
	}

	return db
}

func createTableSession(db *sql.DB) {
	createSessionTableQuery := `
	CREATE TABLE IF NOT EXISTS session (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		start_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		title TEXT,
		tags TEXT,
		duration INTEGER,
		status INTEGER CHECK(status IN (1,2,3))
	);
	`

	_, err := db.Exec(createSessionTableQuery)
	if err != nil {
		log.Fatal("An error occurred while creating DB connection", err)
	}
}
