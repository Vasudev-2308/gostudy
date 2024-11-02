package Validator

import (
	"net/http"

	"github.com/Vasudev-2308/gostudy/intenal/types"
	"github.com/Vasudev-2308/gostudy/intenal/utils/response_util"
	"github.com/go-playground/validator/v10"
)

func StudentValidator(student *types.Student, response http.ResponseWriter) {
	if err := validator.New().Struct(student); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		response_util.WriteToJson(response, http.StatusBadRequest, response_util.ValidationError(validateErrors))
		return
	}
}

func TeacherValidator(teacher *types.Teacher, response http.ResponseWriter) {
	if err := validator.New().Struct(teacher); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		response_util.WriteToJson(response, http.StatusBadRequest, response_util.ValidationError(validateErrors))
		return
	}
}
