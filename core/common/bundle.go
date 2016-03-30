// Author hoenig

package common

import (
	"encoding/json"
	"fmt"
	"io"
)

type Bundle struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Owner   string `json:"owner"`
	Comment string `json:"comment"`
}

func UnpackBundle(r io.Reader) (Bundle, error) {
	return Bundle{}, nil
}

func ValidateBundle(b Bundle) error {
	if !ValidNameRe.MatchString(b.Name) {
		return fmt.Errorf(
			"bundle.Name is bad: '%s'",
			b.Name,
		)
	}

	if b.Path == "" {
		return fmt.Errorf("bundle has empty path")
	}

	if !ValidOwnerRe.MatchString(b.Owner) {
		return fmt.Errorf(
			"bundle.Owner is bad: '%s'",
			b.Owner,
		)
	}

	return nil
}

func (b Bundle) String() string {
	if bs, err := json.Marshal(&b); err != nil {
		return ""
	} else {
		return string(bs)
	}
}
