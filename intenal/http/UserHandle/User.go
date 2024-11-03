package User

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	Validator "github.com/Vasudev-2308/gostudy/intenal/http/validator"
	"github.com/Vasudev-2308/gostudy/intenal/models"
	"github.com/Vasudev-2308/gostudy/intenal/storage"
	response_util "github.com/Vasudev-2308/gostudy/intenal/utils/response_util"
)

func GetUser(database storage.Database, table string) http.HandlerFunc {
	table = strings.ToUpper(table)
	return func(response http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id")
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(err))
			return
		}

		slog.Info("Getting a User")
		user, err := database.GetUserDetail(table, intid)
		if err != nil {
			response_util.WriteToJson(response, http.StatusInternalServerError, response_util.GeneralError(err))
			return
		}

		response_util.WriteToJson(response, http.StatusOK, user)

	}
}

func UpdateUser(database storage.Database, tableName string) http.HandlerFunc {
	tableName = strings.ToUpper(tableName)
	return func(response http.ResponseWriter, request *http.Request) {

		id := request.PathValue("id")
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(err))
			return
		}
		var student models.User
		error := json.NewDecoder(request.Body).Decode(&student)
		if errors.Is(error, io.EOF) {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if error != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(error))
			return
		}
		Validator.UserValidator(&student, response)

		user, err := database.UpdateUser(
			student.Name,
			student.Email,
			student.Subject,
			tableName,
			student.Age,
			intid)

		if err != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(err))
			return
		}

		response_util.WriteToJson(response, http.StatusOK, user)
	}
}

func GetUsers(database storage.Database, tableName string) http.HandlerFunc {
	tableName = strings.ToUpper(tableName)
	return func(response http.ResponseWriter, request *http.Request) {
		users, error := database.GetAllUsers(tableName)
		if error != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(error))
			return
		}
		response_util.WriteToJson(response, http.StatusOK, users)
	}
}

func AddUser(database storage.Database, tableName string) http.HandlerFunc {
	tableName = strings.ToUpper(tableName)
	return func(response http.ResponseWriter, request *http.Request) {
		var newStudent models.User

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
