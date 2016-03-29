// Author hoenig

package agent

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/state"
)

func apiServer(
	address string,
	masters config.Masters,
	tclient *torrent.Client) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(masters, tclient),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Second,
	}
}

func router(masters config.Masters, tclient *torrent.Client) *mux.Router {
	r := mux.NewRouter()
	a := &API{masters: masters, tclient: tclient}
	r.HandleFunc("/v1/create", a.Create).Methods("POST")
	return r
}

type API struct {
	masters config.Masters
	tclient *torrent.Client
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	var bundle state.Bundle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bundle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := state.ValidateBundle(bundle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("/v1/create with name:", bundle.Name, "owner:", bundle.Owner)

	if err := a.create(bundle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("created torrent!\n"))
}

func (a *API) create(bundle state.Bundle) error {
	minfo, err := state.Torrentify(a.masters, bundle, 4)
	if err != nil {
		return err
	}

	// lets check out the metainfo
	log.Println("made torrent ...")
	log.Println("info.Comment", minfo.Comment)
	log.Println("info.CreatedBy", minfo.CreatedBy)
	log.Println("info.CreationDate", minfo.CreationDate)
	return nil
}
