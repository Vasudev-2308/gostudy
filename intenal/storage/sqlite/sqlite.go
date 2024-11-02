package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	"github.com/Vasudev-2308/gostudy/intenal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func NewStudent(cfg *config.Config) (*Sqlite, error) {
	dbInstance, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS STUDENTS (
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

func NewTeacher(cfg *config.Config) (*Sqlite, error) {
	dbInstance, err := sql.Open("sqlite3", cfg.StoragePath)
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

func (db *Sqlite) CreateUser(name string, email string, age int, subject string, tableName string) (int64, error) {

	queryString := fmt.Sprintf("INSERT INTO %s (name, email, age, subject) VALUES (?, ?, ?, ?)", tableName)
	query, err := db.Db.Prepare(
		queryString)

	if err != nil {
		return 0, err
	}

	defer query.Close()
	rslt, err := query.Exec(name, email, age, subject)

	if err != nil {
		return 0, err
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (db *Sqlite) GetUserDetail(tableName string, id int64) (types.User, error) {

	stud := types.User{}
	return stud, nil

}
