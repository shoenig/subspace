// Author hoenig

package stream

import (
	"encoding/json"
	"io"
)

// Bundle represents the publication of a new generation of content on a Stream.
type Bundle struct {
	Info
	Path    string `json:"path"`
	Comment string `json:"comment"`
}

// UnpackBundle unpacks a json representation of a Bundle given an io.Reader.
func UnpackBundle(r io.Reader) (Bundle, error) {
	decoder := json.NewDecoder(r)
	var bundle Bundle
	if err := decoder.Decode(&bundle); err != nil {
		return Bundle{}, err
	}
	if err := bundle.valid(); err != nil {
		return Bundle{}, err
	}
	return bundle, nil
}

// JSON returns the json representation of b.
func (b Bundle) JSON() (string, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
