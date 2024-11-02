package router

import (
	"log/slog"
	"net/http"

	"github.com/Vasudev-2308/gostudy/intenal/config"
	User "github.com/Vasudev-2308/gostudy/intenal/http/UserHandle"
	"github.com/Vasudev-2308/gostudy/intenal/storage/sqlite"
)

func initRoutes(router *http.ServeMux, cfg config.Config) {
	student, err1 := sqlite.NewStudent(&cfg)
	teacher, err2 := sqlite.NewTeacher(&cfg)

	if err1 != nil || err2 != nil {
		slog.Info("Not able to Initiate DB", slog.String("%s", err1.Error()), slog.String("%s", err2.Error()))
	}

	slog.Info("Storage Initated", slog.String("env", cfg.StoragePath))
	router.HandleFunc("GET /api/student/{id}", User.GetUser(student, StudentDB))
	router.HandleFunc("POST /api/create-student", User.AddUser(student, StudentDB))
	router.HandleFunc("GET /api/students", User.GetUsers(student, StudentDB))

	router.HandleFunc("GET /api/teacher/{id}", User.GetUser(teacher, TeacherDB))
	router.HandleFunc("POST /api/create-teachers", User.AddUser(teacher, TeacherDB))
	router.HandleFunc("GET /api/teachers", User.GetUsers(teacher, TeacherDB))
}
