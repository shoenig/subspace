// Author hoenig

package master

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/common/stream"
)

func apiServer(address string, store Store) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(store),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func router(store Store) *mux.Router {
	r := mux.NewRouter()
	api := NewAPI(store)
	r.HandleFunc("/v1/streams", api.GetStreams).Methods("GET")
	r.HandleFunc("/v1/streams/create", api.CreateStream).Methods("PUT")
	r.HandleFunc("/v1/packs/new", api.AddPack).Methods("POST")
	return r
}

// API is the api handler for a master.
type API struct {
	store Store
}

// NewAPI creates a new API backed by store.
func NewAPI(store Store) *API {
	return &API{store: store}
}

// GetStreams returns a JSON list of all streams.
func (a *API) GetStreams(w http.ResponseWriter, r *http.Request) {
	println("master get streams handler")
	w.Header().Set("Content-Type", "application/json")

	streams := a.store.GetStreams()

	encoder := json.NewEncoder(w)
	err := encoder.Encode(streams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

// CreateStream is the handler of a master that will actually create a stream.
func (a *API) CreateStream(w http.ResponseWriter, r *http.Request) {
	println("master create stream")

	s, err := stream.UnpackMetadata(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.store.CreateStream(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(201)
}

// AddPack receives a submitted Pack, which will be made available on the
// stream for clients of agents to download via torrent.
func (a *API) AddPack(w http.ResponseWriter, r *http.Request) {
	println("master add pack handler")

	bundle, err := stream.UnpackBundle(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do logic stuff making pack available

	// add a pack of stuff, return new generation number

	a.store.AddPack(bundle)
}
