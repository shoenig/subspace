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
	"github.com/shoenig/subspace/core/common/subscription"
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
	r.HandleFunc("/v1/subscription/create", a.CreateSubscription).Methods("POST")
	// r.HandleFunc("/v1/create", a.Create).Methods("POST")
	return r
}

// API represents the api handlers of an agent.
type API struct {
	masters config.Masters
	mclient *master.Client
	tclient *torrent.Client
}

// CreateSubscription is an endpoint for creating a new subscription.
func (a *API) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	creation, err := subscription.UnpackCreation(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.createSubscription(creation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *API) createSubscription(c subscription.Creation) error {
	log.Println("client create subscription:", c)
	return a.mclient.CreateSubscription(c)
}

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
