// Author hoenig

package state

import "encoding/json"

type Bundle struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Owner   string `json:"owner"`
	Comment string `json:"comment"`
}

func (b Bundle) String() string {
	if bs, err := json.Marshal(&b); err != nil {
		return ""
	} else {
		return string(bs)
	}
}
