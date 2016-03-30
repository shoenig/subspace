// Author hoenig

package stream

import (
	"encoding/json"
	"fmt"
	"io"
)

// Creation is the information required to request the masters to establish a new Subscription
type Creation struct {
    Info
}

// UnpackCreation unpacks a json representation of a creation given an io.Reader.
func UnpackCreation(r io.Reader) (Creation, error) {
	decoder := json.NewDecoder(r)
	var creation Creation
	if err := decoder.Decode(&creation); err != nil {
		return Creation{}, err
	}

	if err := creation.valid(); err != nil {
		return Creation{}, err
	}

	return creation, nil
}



// String returns a nicely formatted string representation of c.
func (c Creation) String() string {
	return fmt.Sprintf("(%s, %s)", c.Name, c.Owner)
}

// JSON returns the json representation of c.
func (c Creation) JSON() (string, error) {
	bs, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
