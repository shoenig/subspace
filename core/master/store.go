// Author hoenig

package master

import (
	"fmt"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/master/state"
)

// A Store is used by subspace-master to persist data about Stream, Bundle, Info
// and the expected state of the world.
type Store interface {
	NewStream(stream.Metadata) error
	ContainsStream(string) bool
	AllStreams() []stream.Metadata

	NewGeneration(stream.Generation) (uint64, error)
}

// RaftStore is an implemntation of Store based on raft.
type RaftStore struct {
	raft *state.MyRaft
}

// NewRaftStore creates a new RaftStore backed by raft.
func NewRaftStore(raft *state.MyRaft) *RaftStore {
	return &RaftStore{
		raft: raft,
	}
}

// NewStream adds stream to the persisted store.
func (s *RaftStore) NewStream(stream stream.Metadata) error {
	if s.ContainsStream(stream.Name) {
		return fmt.Errorf("stream already exists")
	}
	return s.raft.AddStreams(stream)
}

// ContainsStream returns true if a Stream of name already exists.
func (s *RaftStore) ContainsStream(name string) bool {
	return s.raft.ContainsStream(name)
}

// AllStreams returns a list of the strams.
func (s *RaftStore) AllStreams() []stream.Metadata {
	return s.raft.GetStreams()
}

// NewGeneration will add a new generation to the associated stream.
func (s *RaftStore) NewGeneration(gen stream.Generation) (uint64, error) {
	// add a new generation to the associated stream
	return 0, nil
}
