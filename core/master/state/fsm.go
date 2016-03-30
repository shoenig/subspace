// Author hoenig

package state

import (
	"io"
	"log"
	"sync"

	"github.com/hashicorp/raft"
)

type MasterFSM struct {
	lock sync.Mutex

	streams []stream.Stream
}

func (fsm *MasterFSM) Apply(log *raft.Log) interface{} {
	log.Println("fsm apply")
	return nil
}

func (fsm *MasterFSM) Snapshot() (FSMSnapshot, error) {
	log.Println("fsm snapshot")
	return nil, nil
}

func (fsm *MasterFSM) Restore(closer io.ReadCloser) error {
	log.Println("fsm restore")
	return nil
}
