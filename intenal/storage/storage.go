package storage

import "github.com/Vasudev-2308/gostudy/intenal/types"

type Database interface {
	CreateUser(name string, email string, age int, subject string, tableName string) (int64, error)
	GetUserDetail(tableName string, id int64) (types.User, error)
}
