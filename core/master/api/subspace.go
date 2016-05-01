// Author hoenig

package api

import (
	"encoding/json"
	"net/http"

	"github.com/shoenig/subspace/core/master/service"
)

// SubspaceAPI serves up endpoints exposing meta information about subspace.
type SubspaceAPI struct {
	config *service.Config
}

// NewSubspaceAPI creates a new SubspaceAPI.
func NewSubspaceAPI(config *service.Config) *SubspaceAPI {
	return &SubspaceAPI{config: config}
}

// Masters returns a list of Master information.
func (a *SubspaceAPI) Masters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(a.config.MasterPeers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
