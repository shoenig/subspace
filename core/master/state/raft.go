// Author hoenig

package state

import (
	// "time"

	"encoding/json"
	"time"

	"github.com/hashicorp/raft"
	"github.com/shoenig/subspace/core/common/stream"
)

const (
	timeout = 10 * time.Second // todo, tuneable?
)

// MyRaft is a wrapper around raft for storing the expected state of things.
type MyRaft struct {
	raft *raft.Raft
}

// NewMyRaft creates a new store.
func NewMyRaft() (*MyRaft, error) {
	// rconfig := &raft.Config{
	// 	HeartbeatTimeout:           1 * time.Second,
	// 	ElectionTimeout:            1 * time.Second,
	// 	CommitTimeout:              50 * time.Millisecond,
	// 	MaxAppendEntries:           64,
	// 	ShutdownOnRemove:           true,
	// 	DisableBootstrapAfterElect: true,
	// 	TrailingLogs:               10240,
	// 	SnapshotInterval:           120 * time.Second,
	// 	SnapshotThreshold:          8192,
	// 	EnableSingleNode:           false,
	// 	LeaderLeaseTimeout:         500 * time.Millisecond,
	// 	StartAsLeader:              false,
	// }
	// raft, err := NewRaft(rconfig)
	// if err != nil {
	// 	return nil, err
	// }
	return &MyRaft{
	// raft: raft,
	}, nil
}

func (r *MyRaft) apply(action Action) error {
	bs, err := json.Marshal(action)
	if err != nil {
		return err
	}
	f := r.raft.Apply(bs, timeout)
	return f.Error()
}

// AddStreams to the raft.
func (r *MyRaft) AddStreams(streams ...stream.Stream) error {
	return r.apply(Action{
		Command: AddStreams,
		Streams: streams,
	})
}

// DeleteStreams from the raft.
func (r *MyRaft) DeleteStreams(streams ...stream.Stream) error {
	return r.apply(Action{
		Command: DeleteStreams,
		Streams: streams,
	})
}
