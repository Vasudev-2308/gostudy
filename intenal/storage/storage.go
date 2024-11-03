package storage

import "github.com/Vasudev-2308/gostudy/intenal/models"

type Database interface {
	CreateUser(name string, email string, age int, subject string, tableName string) (int64, error)
	GetUserDetail(tableName string, id int64) (models.User, error)
	GetAllUsers(tableName string) ([]models.User, error)
	UpdateUser(name string, email string, subject string, tableName string, age int, id int64) (models.User, error)
	DeleteUser(tableName string, id int64) (bool, error)
}
