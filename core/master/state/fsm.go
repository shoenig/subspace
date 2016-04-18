// Author hoenig

package state

import (
	"encoding/json"
	"io"
	"log"
	"sync"

	"github.com/hashicorp/raft"
	"github.com/shoenig/subspace/core/common/stream"
)

// MyFSM is an implementation of raft.FSM for creating consensus
// around the state of
// - streams available
// - (todo: other things?)
type MyFSM struct {
	lock sync.RWMutex

	// -- the ultimate source of truth --
	streams map[string]stream.Metadata          // stream.Name -> stream
	packs   map[string]map[uint64]stream.Bundle // stream.Name -> bundles (by generation)
}

// NewMyFSM creates a new MyFSM.
func NewMyFSM() *MyFSM {
	return &MyFSM{
		streams: make(map[string]stream.Metadata),
	}
}

// AddStreams to consensus storage via the FSM.
func (fsm *MyFSM) AddStreams(streams ...stream.Metadata) {
	fsm.lock.Lock()
	defer fsm.lock.Unlock()

	fsm.addStreams(streams...)
}

// fsm.lock must be held
func (fsm *MyFSM) addStreams(streams ...stream.Metadata) {
	for _, stream := range streams {
		fsm.streams[stream.Name] = stream
	}
}

// DeleteStreams from consensus storage via FSM.
// It is acceptable for each Stream to only have Name set.
func (fsm *MyFSM) DeleteStreams(streams ...stream.Metadata) {
	fsm.lock.Lock()
	defer fsm.lock.Unlock()

	for _, stream := range streams {
		delete(fsm.streams, stream.Name)
	}
}

// CopyStreams from consensus storage via FSM.
func (fsm *MyFSM) CopyStreams() []stream.Metadata {
	streams := make([]stream.Metadata, 0, len(fsm.streams))
	fsm.lock.RLock()
	defer fsm.lock.RUnlock()

	for _, stream := range fsm.streams {
		streams = append(streams, stream)
	}
	return streams
}

// ContainsStream returns true if the FSM contains a stream of the given name.
func (fsm *MyFSM) ContainsStream(name string) bool {
	fsm.lock.RLock()
	defer fsm.lock.RUnlock()

	_, exists := fsm.streams[name]
	return exists
}

// Do an Action on the FSM.
func (fsm *MyFSM) Do(action Action) {
	switch action.Command {
	case AddStreams:
		fsm.AddStreams(action.Streams...)
	case DeleteStreams:
		fsm.DeleteStreams(action.Streams...)
	default:
		panic("there is a fatal bug in this program")
	}
}

// -- FSM interface below this line --

// Apply will apply log to the FSM.
func (fsm *MyFSM) Apply(entry *raft.Log) interface{} {
	log.Println("fsm apply")
	var action Action
	if err := json.Unmarshal(entry.Data, &action); err != nil {
		panic("fatal error unpacking an action:" + err.Error())
	}
	fsm.Do(action)
	return nil // why does this return interface{}?
}

// Snapshot will take a snapshot of the FSM.
func (fsm *MyFSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("fsm snapshot")
	return &MySnapshot{
		streams: fsm.CopyStreams(),
	}, nil
}

// Restore will restore the state of the FSM.
func (fsm *MyFSM) Restore(snapshot io.ReadCloser) error {
	log.Println("fsm restore")
	decoder := json.NewDecoder(snapshot)
	var streams []stream.Metadata
	if err := decoder.Decode(&streams); err != nil {
		snapshot.Close() // close no matter what
		return err
	}

	// reset the whole FSM with the snapshot
	fsm.lock.Lock()
	fsm.streams = make(map[string]stream.Metadata, len(streams))
	fsm.addStreams(streams...)
	fsm.lock.Unlock()

	return snapshot.Close()
}
