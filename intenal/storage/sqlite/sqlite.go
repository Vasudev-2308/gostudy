package sqlite

import (
	"database/sql"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	dbInstance, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS STUDENTS (
			ID INTEGER PRIMARY KEY AUTOINCREMENT, 
			NAME TEXT, 
			AGE INTEGER, 
			EMAIL TEXT
		)
	`)

	if err != nil {
		return nil, err
	}

	_, err = dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS TEACHERS (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			NAME TEXT,
			AGE INTEGER,
			EMAIL TEXT,
			SUBJECT TEXT
		)
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: dbInstance,
	}, nil

}

func (db *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	query, err := db.Db.Prepare(
		"INSERT INTO STUDENTS (name, email, age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer query.Close()
	rslt, err := query.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (db *Sqlite) CreateTeacher(name string, email string, age int, subject string) (int64, error) {
	query, err := db.Db.Prepare(
		"INSERT INTO TEACHERS (name, email, subject, age) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer query.Close()
	rslt, err := query.Exec(name, email, subject, age)

	if err != nil {
		return 0, err
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}
