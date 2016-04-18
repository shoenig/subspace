// Author hoenig

package stream

import (
	"fmt"
	"regexp"
)

// Rules for validating name and owner of a stream.
var (
	ValidNameRe  = regexp.MustCompile("^[a-z0-9]+[a-z_-]*$")
	ValidOwnerRe = regexp.MustCompile("^[a-z0-9]+$")
)

// ValidName returns an error if name does not comply with the validation regex
// for stream names, as determined by ValidNameRe.
func ValidName(name string) error {
	if !ValidNameRe.MatchString(name) {
		return fmt.Errorf("stream name is bad: '%s'", name)
	}
	return nil
}

// ValidOwner returns an error if owner does not comply with the validation regex
// for stream owners, as determined by ValidOwnerRe.
func ValidOwner(owner string) error {
	if !ValidOwnerRe.MatchString(owner) {
		return fmt.Errorf("owner name is bad: '%s'", owner)
	}
	return nil
}
