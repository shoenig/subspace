// Author hoenig

package agent

import (
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master"
	"github.com/shoenig/subspace/core/common/stream"
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
	// r.HandleFunc("/v1/create", a.Create).Methods("POST")
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
	creation, err := stream.UnpackCreation(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.CreateStream(creation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *API) createStream(c stream.Creation) error {
	log.Println("client create stream:", c)
	return a.mclient.CreateStream(c)
}




// an old example of creating a torrent

// func (a *API) Create(w http.ResponseWriter, r *http.Request) {
// 	var bundle common.Bundle
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&bundle); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := common.ValidateBundle(bundle); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	log.Println("/v1/create with name:", bundle.Name, "owner:", bundle.Owner)

// 	minfo, err := a.create(bundle)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	s := fmt.Sprintf("created torrent: %s\n", minfo.Info.Name)
// 	w.Write([]byte(s))
// }

// func (a *API) create(bundle common.Bundle) (*metainfo.MetaInfo, error) {
// 	minfo, err := common.Torrentify(a.masters, bundle, 4)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// lets check out the metainfo
// 	log.Println("made torrent ...")
// 	log.Println("info.Comment", minfo.Comment)
// 	log.Println("info.CreatedBy", minfo.CreatedBy)
// 	log.Println("info.CreationDate", minfo.CreationDate)
// 	return minfo, nil
// }
