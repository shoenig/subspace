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
	a := &api{masters: masters, tclient: tclient}
	r.HandleFunc("/v1/create", a.create).Methods("POST")
	return r
}

type api struct {
	masters config.Masters
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

	// create torrent of bundle
	if minfo, err := state.Torrentify(a.masters, bundle, 4); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// lets check out the metainfo
		log.Println("made torrent ...")
		log.Println("info.Comment", minfo.Comment)
		log.Println("info.CreatedBy", minfo.CreatedBy)
		log.Println("info.CreationDate", minfo.CreationDate)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("created torrent!"))
	}
}
