// Author hoenig

package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/master/state"
)

// Server returns an http.Server serving up the API.
func Server(address string, store state.Store) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(store),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func router(store state.Store) *mux.Router {
	r := mux.NewRouter()
	api := NewAPI(store)
	r.HandleFunc("/v1/streams", api.AllStreams).Methods("GET")
	r.HandleFunc("/v1/streams/create", api.NewStream).Methods("PUT")
	r.HandleFunc("/v1/streams/publish", api.PublishGeneration).Methods("POST")
	return r
}
