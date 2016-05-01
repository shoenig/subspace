// Author hoenig

package api

import (
	"encoding/json"
	"net/http"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/master/state"
)

// StreamsAPI is the api handler for managing streams.
type StreamsAPI struct {
	store state.Store
}

// NewStreamsAPI creates a new API backed by store.
func NewStreamsAPI(store state.Store) *StreamsAPI {
	return &StreamsAPI{store: store}
}

// AllStreams returns a JSON list of all streams.
func (a *StreamsAPI) AllStreams(w http.ResponseWriter, r *http.Request) {
	println("master get streams handler")
	w.Header().Set("Content-Type", "application/json")
	streams := a.store.AllStreams()

	encoder := json.NewEncoder(w)
	err := encoder.Encode(streams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// NewStream is the handler of a master that will actually create a stream.
func (a *StreamsAPI) NewStream(w http.ResponseWriter, r *http.Request) {
	println("master create stream")

	s, err := stream.UnpackMetadata(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.store.NewStream(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(201)
}

// PublishGeneration receives a submitted generation, which will be made available on the
// stream for clients of agents to download via torrent.
func (a *StreamsAPI) PublishGeneration(w http.ResponseWriter, r *http.Request) {
	println("master new generation published handler")

	gen, err := stream.UnpackGeneration(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do logic stuff making generation available

	// add a generation full of stuff, return new generation number

	a.store.NewGeneration(gen)
}
