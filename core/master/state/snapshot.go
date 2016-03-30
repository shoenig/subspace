// Author hoenig

package state

import (
	"log"

	"github.com/hashicorp/raft"
)

// masterSnapshot is used to provide a snapshot of the curren state
// in a way that can be accessed concurrently with operations that may
// modify the live state.
type masterSnapshot struct {
}

func (sn *masterSnapshot) Persist(sink raft.SnapshotSink) error {
	log.Println("snapshot persist")
	return nil
}

func (sn *masterSnapshot) Release() {
	log.Println("snapshot release")
}
