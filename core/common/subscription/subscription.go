// Author hoenig

package subscription

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/shoenig/subspace/core/common"
)

// Creation is the information required to request the masters to establish a new Subscription
type Creation struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
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

func (c Creation) valid() error {
	if !common.ValidNameRe.MatchString(c.Name) {
		return fmt.Errorf("subscription creation name is bad: '%s'", c.Name)
	}

	if !common.ValidOwnerRe.MatchString(c.Owner) {
		return fmt.Errorf("subscription creation owner is bad: '%s'", c.Owner)
	}

	return nil
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

// Publication is the information sent from an agent that created a torrent
// type Publication struct {
// 	Name string `json:"subscription"`
// }
