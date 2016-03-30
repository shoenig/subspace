// Author hoenig

package state

import (
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
func (fsm *MyFSM) DeleteStreams(names ...string) {
	fsm.lock.Lock()
	for _, name := range names {
		delete(fsm.streams, name)
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

// -- FSM interface below this line --

// Apply will apply log to the FSM.
func (fsm *MyFSM) Apply(entry *raft.Log) interface{} {
	log.Println("fsm apply")
	return nil
}

// Snapshot will take a snapshot of the FSM.
func (fsm *MyFSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("fsm snapshot")
	return nil, nil
}

// Restore will restore the state of the FSM.
func (fsm *MyFSM) Restore(closer io.ReadCloser) error {
	log.Println("fsm restore")
	return nil
}
