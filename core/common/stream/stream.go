// Author hoenig

package stream

import (
	"fmt"

	"github.com/shoenig/subspace/core/common"
)

// Info represents the fundamental identification of a Stream
type Info struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (i Info) valid() error {
	if !common.ValidNameRe.MatchString(i.Name) {
		return fmt.Errorf("stream name is bad: '%s'", i.Name)
	}

	if !common.ValidOwnerRe.MatchString(i.Owner) {
		return fmt.Errorf("stream owner is bad: '%s'", i.Owner)
	}

	return nil
}

// A Stream represents the flow of things that can be downloaded.
type Stream struct {
	Info
}

// NewStream creates a new stream with the given name and owner.
func New(name, owner string) Stream {
	return Stream{
		Info: Info{
			Name:  name,
			Owner: owner,
		},
	}
}
