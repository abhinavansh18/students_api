package student

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/abhinavansh18/students_api/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		json.NewDecoder(r.Body).Decode(&student)

		slog.Info(("Creating a student"))
		w.Write([]byte("Welcome to students api"))

	}
}
