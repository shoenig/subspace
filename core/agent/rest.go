// Author hoenig

package agent

import (
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/v1/create/{name}", a.newBundle)
	return r
}

type api struct {
	tclient *torrent.Client
}

func (a *api) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("/v1/create with name=", vars["name"])
}
