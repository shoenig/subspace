// Author hoenig

package state

import (
	"encoding/json"
	"os"
	"time"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/shoenig/subspace/core/common/stream"
)

const (
	timeout = 10 * time.Second // todo, tuneable?
)

// MyRaft is a wrapper around raft for storing the expected state of things.
type MyRaft struct {
	raft *raft.Raft

	// hold on to these so we can close them on graceful shutdown
	transport *raft.NetworkTransport
	boltstore *raftboltdb.BoltStore
}

// NewMyRaft creates a new store.
func NewMyRaft() (*MyRaft, error) {
	rconfig := &raft.Config{
		HeartbeatTimeout:           1 * time.Second,
		ElectionTimeout:            1 * time.Second,
		CommitTimeout:              50 * time.Millisecond,
		MaxAppendEntries:           64,
		ShutdownOnRemove:           true,
		DisableBootstrapAfterElect: true,
		TrailingLogs:               10240,
		SnapshotInterval:           120 * time.Second,
		SnapshotThreshold:          8192,
		EnableSingleNode:           false,
		LeaderLeaseTimeout:         500 * time.Millisecond,
		StartAsLeader:              false,
	}

	boltstore, err := raftboltdb.NewBoltStore("/todo")
	if err != nil {
		return nil, err
	}

	filestore, err := raft.NewFileSnapshotStore("/todo", 1, os.Stdout)
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(
		"0.0.0.0:0",   // bind address
		nil,           // advertise address
		0,             // maxPool (unused)
		5*time.Second, // timeout
		os.Stdout,     // debug log file
	)
	if err != nil {
		return nil, err
	}

	peerstore := raft.NewJSONPeers("/todo", transport)

	raft, err := raft.NewRaft(
		rconfig,    // raft config
		NewMyFSM(), // fsm implementation
		boltstore,  // raft log store
		boltstore,  // raft stable store
		filestore,  // snapshot store
		peerstore,  // peer store (human editable)
		transport,  // the internet
	)
	if err != nil {
		return nil, err
	}

	return &MyRaft{
		raft:      raft,
		boltstore: boltstore,
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

// Close will attempt to gracefully stop the raft. Although we
// must be totally resiliant to failure at anytime, a controlled
// shutdown allows remaining raft members to recover more quickley.
func (r *MyRaft) Close() error {
	if err := r.transport.Close(); err != nil {
		r.boltstore.Close() // try to close the boltstore anyway
		return err
	}

	if err := r.boltstore.Close(); err != nil {
		return err
	}
	return nil
}
