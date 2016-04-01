// Author hoenig

package state

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/shoenig/subspace/core/common/stream"
)

const (
	timeout = 10 * time.Second // todo, tuneable?
)

// Config allows for fine-tuning how hashicorp/raft will operate.
type Config struct {
	BindAddress string `json:"bind.address"`
	DataDir     string `json:"data.dir"`
	SingleMode  bool   `json:"single.master.mode"`
}

// MyRaft is a wrapper around raft for storing the expected state of things.
type MyRaft struct {
	raft *raft.Raft

	// hold on to these so we can close them on graceful shutdown
	transport *raft.NetworkTransport
	boltstore *raftboltdb.BoltStore
}

// NewMyRaft creates a new store.
func NewMyRaft(leader bool, rcfg Config) (*MyRaft, error) {
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
		EnableSingleNode:           rcfg.SingleMode,
		LeaderLeaseTimeout:         500 * time.Millisecond,
		StartAsLeader:              leader,
	}

	boltPath := filepath.Join(rcfg.DataDir, "boltdb")
	if err := touchDB(boltPath); err != nil {
		return nil, err
	}
	boltstore, err := raftboltdb.NewBoltStore(boltPath)
	if err != nil {
		return nil, err
	}

	// filestore automatically creates directory called snapshots
	filestore, err := raft.NewFileSnapshotStore(rcfg.DataDir, 1, os.Stdout)
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(
		rcfg.BindAddress, // bind address
		nil,              // advertise address
		0,                // maxPool (unused)
		5*time.Second,    // timeout
		os.Stdout,        // debug log file
	)
	if err != nil {
		return nil, err
	}

	peerPath := filepath.Join(rcfg.DataDir, "peers")
	peerstore := raft.NewJSONPeers(peerPath, transport)

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

// filesystems are the best
func touchDB(boltpath string) error {
	base := filepath.Dir(boltpath)
	if err := os.MkdirAll(base, 0700); err != nil {
		return err
	}

	// try to touch the file
	touch := os.Chtimes(boltpath, time.Now(), time.Now())
	if touch != nil && strings.Contains(touch.Error(), "no such file") {
		log.Println("will create new boltdb at", boltpath)

		// create + fsync + close the boltdb file
		f, err := os.OpenFile(boltpath, os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		if err := f.Sync(); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}

		// fsync + close the parent directory
		d, err := os.Open(base)
		if err != nil {
			return err
		}
		if err := d.Sync(); err != nil {
			return err
		}
		return d.Close()

	} else if touch != nil {
		// some other error occured
		return touch
	}

	log.Println("will use existing boltdb at", boltpath)
	return nil
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
