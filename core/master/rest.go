// Author hoenig

package master

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoenig/subspace/core/state/subscription"
)

func apiServer(address string) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func router() *mux.Router {
	r := mux.NewRouter()
	a := &API{}
	r.HandleFunc("/v1/subscription/create", a.CreateSubscription).Methods("POST")
	return r
}

// API is the api handler for a master.
type API struct {
}

// CreateSubscription is the handler of a master that will actually create a subscription.
func (a *API) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	creation, err := subscription.UnpackCreation(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.createSubscription(creation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	msg := fmt.Sprintf("create subscription %v", creation)
	w.Write([]byte(msg))
}

func (a *API) createSubscription(c subscription.Creation) error {
	log.Println("master will create subscription:", c)
	return nil
}
