package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func Init(dbPath string) error {
	var err error
	once.Do(func() {
		dir := filepath.Dir(dbPath)
		if err = os.MkdirAll(dir, 0755); err != nil {
			return
		}
		connStr := dbPath + "?_journey_mode=WAL"
		db, err = sql.Open("sqlite3", connStr)
		if err != nil {
			return
		}
		db.SetMaxOpenConns(1)
	})
	return err
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
