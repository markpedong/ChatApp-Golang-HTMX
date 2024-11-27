package routes

import (
	"encoding/json"
	"net/http"
)

type ErrJSON struct {
	Message string `json:"message"`
}

func CreateRoutes(s *http.ServeMux) {
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		j, err := json.MarshalIndent(ErrJSON{"Not supported yet."}, "\n", "    ")
		if err != nil {
			return
		}

		http.Error(w, string(j), http.StatusBadRequest)
	})
}
