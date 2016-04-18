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
	CreateStream(stream.Metadata) error
	ContainsStream(string) bool
	GetStreams() []stream.Metadata

	AddPack(stream.Bundle) (uint64, error)
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

// CreateStream adds stream to the persisted store.
func (s *RaftStore) CreateStream(stream stream.Metadata) error {
	if s.ContainsStream(stream.Name) {
		return fmt.Errorf("stream already exists")
	}
	return s.raft.AddStreams(stream)
}

// ContainsStream returns true if a Stream of name already exists.
func (s *RaftStore) ContainsStream(name string) bool {
	return s.raft.ContainsStream(name)
}

// GetStreams returns a list of the strams.
func (s *RaftStore) GetStreams() []stream.Metadata {
	return s.raft.GetStreams()
}

func (s *RaftStore) AddPack(pack stream.Bundle) (uint64, error) {
	// add pack
	return 0, nil
}
