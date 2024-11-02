package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	"github.com/Vasudev-2308/gostudy/intenal/models"
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

func (db *Sqlite) GetUserDetail(tableName string, id int64) (models.User, error) {

	queryStmt := fmt.Sprintf("SELECT * FROM %s WHERE ID = ? LIMIT 1 ", tableName)
	query, err := db.Db.Prepare(queryStmt)

	if err != nil {
		return models.User{}, err
	}

	defer query.Close()
	var user models.User

	err = query.QueryRow(id).Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no student found with id: %d", id)
		}
		return models.User{}, fmt.Errorf("query error %w", err)
	}

	return user, nil
}

func (db *Sqlite) GetAllUsers(tableName string) ([]models.User, error) {

	queryStmt := fmt.Sprintf("SELECT ID, NAME, AGE, EMAIL, SUBJECT FROM %s ", tableName)
	query, err := db.Db.Prepare(queryStmt)

	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	var students []models.User
	for rows.Next() {
		var student models.User
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Subject)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
