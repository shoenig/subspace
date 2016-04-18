// Author hoenig

package agent

import (
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master"
)

func apiServer(
	address string,
	masters config.Masters,
	mclient *master.Client,
	tclient *torrent.Client) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(masters, mclient, tclient),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Second,
	}
}

func router(masters config.Masters, mclient *master.Client, tclient *torrent.Client) *mux.Router {
	r := mux.NewRouter()
	a := &API{masters: masters, mclient: mclient, tclient: tclient}
	r.HandleFunc("/v1/stream/create", a.CreateStream).Methods("POST")
	return r
}

// API represents the api handlers of an agent.
type API struct {
	masters config.Masters
	mclient *master.Client
	tclient *torrent.Client
}

// CreateStream is an endpoint for creating a new stream.
func (a *API) CreateStream(w http.ResponseWriter, r *http.Request) {
	stream, err := stream.UnpackMetadata(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.createStream(stream); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *API) createStream(c stream.Metadata) error {
	log.Println("client create stream:", c)
	return a.mclient.CreateStream(c)
}

// Publish a new Generation to a stream.
func (a *API) Publish(w http.ResponseWriter, r *http.Request) {
	gen, err := stream.UnpackGeneration(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.publish(gen); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// given a Generation, Torrentify the content creating a MagnetURI, and
// POST the information to a master.
func (a *API) publish(gen stream.Generation) error {
	log.Println("publishing new generation:", gen)
	/*
		mi, err := common.Torrentify(a.masters, b, 4)
		if err != nil {
			return err
		}
		magnet := mi.Magnet()
		thing := stream.NewThing(b, magnet.String())
		return a.mclient.Publish(thing)
	*/
	return nil
}
