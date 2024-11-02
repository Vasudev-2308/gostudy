package User

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	Validator "github.com/Vasudev-2308/gostudy/intenal/http/validator"
	"github.com/Vasudev-2308/gostudy/intenal/storage"
	"github.com/Vasudev-2308/gostudy/intenal/types"
	response_util "github.com/Vasudev-2308/gostudy/intenal/utils/response_util"
)

func GetUser(database storage.Database, table string) http.HandlerFunc {
	table = strings.ToUpper(table)
	return func(response http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id")
		slog.Info(id)
		if table == "STUDENT" {
			slog.Info("Getting a Student")
			user, err := database.GetUserDetail(table, 2)
			if err != nil {
				slog.Info(err.Error())
			}
			fmt.Println(user)
		}
		if table == "TEACHER" {
			slog.Info("Getting a Student")
			user, err := database.GetUserDetail(table, 2)
			if err != nil {
				slog.Info(err.Error())
			}
			fmt.Println(user)
		}
	}
}

func AddUser(database storage.Database, tableName string) http.HandlerFunc {
	tableName = strings.ToUpper(tableName)
	return func(response http.ResponseWriter, request *http.Request) {
		var newStudent types.User

		error := json.NewDecoder(request.Body).Decode(&newStudent)

		if errors.Is(error, io.EOF) {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if error != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(error))
			return
		}
		Validator.UserValidator(&newStudent, response)

		id, err := database.CreateUser(
			newStudent.Name,
			newStudent.Email,
			newStudent.Age,
			newStudent.Subject,
			tableName)

		slog.Info("User Created", slog.String("id : %s", string(id)))
		if err != nil {
			response_util.WriteToJson(response, http.StatusInternalServerError, err)
			return
		}

		response_util.WriteToJson(response, http.StatusCreated, map[string]int64{"id": id})
	}
}