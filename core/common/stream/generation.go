// Author hoenig

package stream

import (
	"encoding/json"
	"fmt"
	"io"
)

// Generation represents the publication of a new generation of content to a stream.
type Generation struct {
	Stream    string `json:"stream"`
	Path      string `json:"path"`
	Comment   string `json:"comment"`
	MagnetURI string `json:"magnet"`
}

// UnpackGeneration unpacks a json representation of a Generation.
func UnpackGeneration(r io.Reader) (Generation, error) {
	decoder := json.NewDecoder(r)
	var gen Generation
	if err := decoder.Decode(&gen); err != nil {
		return Generation{}, err
	}
	if err := gen.valid(); err != nil {
		return Generation{}, err
	}
	return gen, nil
}

// JSON returns the json representation of b.
func (b Generation) JSON() (string, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (b Generation) valid() error {
	if err := ValidName(b.Stream); err != nil {
		return err
	}
	if b.Path == "" {
		return fmt.Errorf("generation cannot represent empty file path")
	}
	return nil
}
