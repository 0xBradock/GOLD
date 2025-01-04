package user

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/0xBradock/go-srvr/internal/utils"
)

func HandleUser(l *slog.Logger, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			utils.ErrorResponse(w, "missing header for 'application/json'", http.StatusBadRequest)
		}

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.ErrorResponse(w, "failed to read request body", http.StatusBadRequest)
			return
		}

		u := &User{}
		if err := json.Unmarshal(reqBody, u); err != nil {
			msg := "request do not contain required fields"
			l.Error(msg, "body", reqBody)
			utils.ErrorResponse(w, msg, http.StatusBadRequest)
		}

		if err := store.Store(u); err != nil {
			msg := "failed to store given user"
			l.Error(msg, "user", u)
			utils.ErrorResponse(w, msg, http.StatusInternalServerError)
		}

		l.Info("successfully stored given user")
		utils.SuccessResponse(w, http.StatusCreated)
	}
}
