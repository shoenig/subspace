// Author hoenig

package state

import "github.com/shoenig/subspace/core/common/stream"

// Command indicates what is to be done to the FSM.
type Command int

// Commands that indicate what to do to the FSM.
// Never change the order (maintain backwards compatility).
const (
	AddStreams Command = iota
	DeleteStreams
)

// Action is something to be applied to the FSM.
type Action struct {
	Command Command         `json:"command"`
	Streams []stream.Stream `json:"streams"`
}
