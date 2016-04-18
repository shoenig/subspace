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
	node string,
	address string,
	masters config.Masters,
	mclient *master.Client) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(node, masters, mclient),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Second,
	}
}

func router(node string, masters config.Masters, mclient *master.Client) *mux.Router {
	r := mux.NewRouter()
	a := &API{
		node:    node,
		masters: masters,
		mclient: mclient,
	}
	r.HandleFunc("/v1/streams/create", a.CreateStream).Methods("PUT")
	r.HandleFunc("/v1/streams/publish", a.PublishGeneration).Methods("POST")
	return r
}

// API represents the api handlers of an agent.
type API struct {
	node    string
	masters config.Masters

	mclient *master.Client
	tclient *torrent.Client
}

// CreateStream is an endpoint for creating a new stream, available on every
// agent so that clients do not need to know anything about the masters.
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

// PublishGeneration singals that a new generation of data on a stream is
// available for torrenting.
func (a *API) PublishGeneration(w http.ResponseWriter, r *http.Request) {
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

	// going to need metadata of gen.stream
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
