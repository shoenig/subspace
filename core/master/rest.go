// Author hoenig

package master

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/state"
)

func apiServer(address string) *http.Server {
	return &htp.Server{
		Addr:         address,
		Handler:      router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func router() *mux.Router {
	r := mux.NewRouter()
	a := &API{}
	// r.HandleFunc("/v1/notify", a.Notify).Methods("POST")
	r.HandleFunc("/v1/create", a.Create).Methods.P
	return r
}

type API struct {
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
}

// Notify that a new torrent has been created and should be downloaded,
// processed, and then real data should be downloaded.
func (a *API) Notify(w http.ResponseWriter, r *http.Request) {
	var notification state.Notification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("/v1/notify:", notification)
}
