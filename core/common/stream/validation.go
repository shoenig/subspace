// Author hoenig

package stream

import "regexp"

var (
	ValidNameRe  = regexp.MustCompile("^[a-z]+[a-z_-]*$")
	ValidOwnerRe = regexp.MustCompile("^[a-z]+$")
)
