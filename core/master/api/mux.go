// Author hoenig

package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/master/service"
	"github.com/shoenig/subspace/core/master/state"
)

// Server returns an http.Server serving up the API.
func Server(address string, config *service.Config, store state.Store) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(config, store),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func router(config *service.Config, store state.Store) *mux.Router {
	r := mux.NewRouter()

	ss := NewSubspaceAPI(config)
	r.HandleFunc("/v1/subspace/masters", ss.Masters).Methods("GET")

	api := NewStreamsAPI(store)
	r.HandleFunc("/v1/streams", api.AllStreams).Methods("GET")
	r.HandleFunc("/v1/streams/create", api.NewStream).Methods("PUT")
	r.HandleFunc("/v1/streams/publish", api.PublishGeneration).Methods("POST")
	return r
}
