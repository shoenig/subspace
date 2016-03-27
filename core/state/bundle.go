// Author hoenig

package state

import (
	"encoding/json"
	"fmt"
)

type Bundle struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Owner   string `json:"owner"`
	Comment string `json:"comment"`
}

func Validate(b Bundle) error {
	if b.Name == "" {
		return fmt.Errorf("bundle has empty name")
	}

	if b.Path == "" {
		return fmt.Errorf("bundle has empty path")
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
