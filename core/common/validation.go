// Author hoenig

package common

import "regexp"

var (
	ValidNameRe  = regexp.MustCompile("^[a-z]+[a-z_-]*$")
	ValidOwnerRe = regexp.MustCompile("^[a-z]+$")
)
