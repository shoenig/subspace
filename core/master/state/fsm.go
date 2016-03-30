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

	// the source of truth
	streams map[string]stream.Stream
}

// NewMyFSM creates a new MyFSM.
func NewMyFSM() *MyFSM {
	return &MyFSM{
		streams: make(map[string]stream.Stream),
	}
}

// AddStreams to consensus storage via the FSM.
func (fsm *MyFSM) AddStreams(streams ...stream.Stream) {
	fsm.lock.Lock()
	for _, stream := range streams {
		fsm.streams[stream.Name] = stream
	}
	fsm.lock.Unlock()
}

// DeleteStreams from consensus storage via FSM.
// It is acceptable for each Stream to only have Name set.
func (fsm *MyFSM) DeleteStreams(streams ...stream.Stream) {
	fsm.lock.Lock()
	for _, stream := range streams {
		delete(fsm.streams, stream.Name)
	}
	fsm.lock.Unlock()
}

// CopyStreams from consensus storage via FSM.
func (fsm *MyFSM) CopyStreams() []stream.Stream {
	streams := make([]stream.Stream, 0, len(fsm.streams))
	fsm.lock.RLock()
	for _, stream := range fsm.streams {
		streams = append(streams, stream)
	}
	fsm.lock.RUnlock()
	return streams
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
	var streams []stream.Stream
	if err := decoder.Decode(&streams); err != nil {
		snapshot.Close() // close no matter what
		return err
	}

	// reset the whole FSM with the snapshot
	fsm.lock.Lock()
	fsm.streams = make(map[string]stream.Stream, len(streams))
	fsm.AddStreams(streams...)
	fsm.lock.Unlock()

	return snapshot.Close()
}
