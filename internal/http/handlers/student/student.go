package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/abhinavansh18/students_api/internal/types"
	"github.com/abhinavansh18/students_api/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("Empty Body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
		}
		//request validation
		validate := validator.New()
		if err := validate.Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return

		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
