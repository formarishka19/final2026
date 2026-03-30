package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

const SCHEMA = `CREATE TABLE IF NOT EXISTS scheduler (
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    date    CHAR(8) NOT NULL DEFAULT "",
    title   VARCHAR(128) NOT NULL,
    comment TEXT,
    repeat  VARCHAR(128)
);
CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date);`

var db *sql.DB

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)
	if err == nil {
		db, err = sql.Open("sqlite", dbFile)
		if err != nil {
			log.Fatal("error acces to database %w", err)
		}
	} else {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("creating db file")
			db, err = sql.Open("sqlite", dbFile)
			if err != nil {
				log.Fatal("error acces to database %w", err)
			}
			defer db.Close()
			_, err = db.Exec(SCHEMA)
			if err != nil {
				log.Fatal("error generating db schema %w", err)
			}
		} else {
			log.Fatal("error accessing db file, possible no access %w", err)
		}
	}
	return err
}
