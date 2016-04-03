// Author hoenig

package stream

import "regexp"

var (
	ValidNameRe  = regexp.MustCompile("^[a-z0-9]+[a-z_-]*$")
	ValidOwnerRe = regexp.MustCompile("^[a-z0-9]+$")
)
