// Author hoenig

package agent

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/state"
)

func apiServer(addr string, tclient *torrent.Client) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router(tclient),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Second,
	}
}

func router(tclient *torrent.Client) *mux.Router {
	r := mux.NewRouter()
	a := &api{tclient: tclient}
	r.HandleFunc("/v1/create", a.create).Methods("POST")
	return r
}

type api struct {
	tclient *torrent.Client
}

func (a *api) create(w http.ResponseWriter, r *http.Request) {
	var bundle state.Bundle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bundle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := state.Validate(bundle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("/v1/create with name:", bundle.Name, "owner:", bundle.Owner)
}
