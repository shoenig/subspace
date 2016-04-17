// Author hoenig

package master

import (
	"fmt"
	"log"
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
	r.HandleFunc("/v1/stream/create", api.CreateStream).Methods("POST")
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

// CreateStream is the handler of a master that will actually create a stream.
func (a *API) CreateStream(w http.ResponseWriter, r *http.Request) {
	println("master create stream")
	creation, err := stream.UnpackCreation(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.createStream(creation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	msg := fmt.Sprintf("create stream %v", creation)
	w.Write([]byte(msg))
}

func (a *API) createStream(c stream.Creation) error {
	log.Println("master will create stream:", c)
	return nil
}
