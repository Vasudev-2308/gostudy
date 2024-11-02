package storage

type Database interface {

	CreateStudent(name string, email string, age int, subject string) (int64, error)
	CreateTeacher(name string, email string, age int, subject string) (int64, error)
}