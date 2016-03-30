// Author hoenig

package state

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/raft"
	"github.com/shoenig/subspace/core/common/stream"
)

// MySnapshot is an FSMSnapshot that is used to provide a snapshot of
// the current state in a way that can be accessed concurrently with
// operations that may modify the live state - it must be safe to invoke
// Persist and Release methods with concurrent calls to FSM.Apply.
type MySnapshot struct {
	// keep our own copy of the data
	// keep in mind, this thing needs to be jsonable of all data
	// (todo, create an addressable wrapper when we add stuff)
	streams []stream.Stream
}

// Persist dumps all necessary state to the WriteCloser sink, and then
// call sink.Close when finished or sink.Cancel on error.
func (sn *MySnapshot) Persist(sink raft.SnapshotSink) error {
	log.Println("snapshot persist")

	bs, err := json.Marshal(sn.streams)
	if err != nil {
		return err
	}

	_, err = sink.Write(bs)
	if err != nil {
		sink.Cancel()
		return err
	}

	return sink.Close()
}

// Release is invoked when we are finished with the snapshot.
func (sn *MySnapshot) Release() {
	log.Println("snapshot release")
	// nothing to do here
}
