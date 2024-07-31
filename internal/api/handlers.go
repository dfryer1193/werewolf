package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (api *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/ping", api.PingHandler).Methods("GET")
}

func (api *API) PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Ping handler called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}
