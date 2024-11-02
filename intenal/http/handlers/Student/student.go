package Student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	Validator "github.com/Vasudev-2308/gostudy/intenal/http/validator"
	"github.com/Vasudev-2308/gostudy/intenal/types"
	response_util "github.com/Vasudev-2308/gostudy/intenal/utils/response_util"
)

func GetStudents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Student"))
	}
}

func CreateStudent() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var newStudent types.Student

		error := json.NewDecoder(request.Body).Decode(&newStudent)

		if errors.Is(error, io.EOF) {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if error != nil {
			response_util.WriteToJson(response, http.StatusBadRequest, response_util.GeneralError(error))
			return
		}
		 Validator.StudentValidator(&newStudent, response)

		response_util.WriteToJson(response, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
