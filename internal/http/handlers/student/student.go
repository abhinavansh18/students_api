package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/abhinavansh18/students_api/internal/types"
	"github.com/abhinavansh18/students_api/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("Empty Body")))
			return
		}

		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
