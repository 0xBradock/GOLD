package health

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func HandleHealth(l *slog.Logger) http.HandlerFunc {
	type response struct {
		Message string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		l.Info("health")
		res := &response{Message: "Ok"}

		b, _ := json.Marshal(res)
		w.Write(b)
	}
}
