// Author hoenig

// Package stream contains struct definitions which represent a flow of
// generational data is to be downloaded via torrent by many clients.
package stream

import (
	"encoding/json"
	"io"
	"time"
)

// Metadata represents the fundamental properties of a Stream.
type Metadata struct {
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Created int64  `json:"created"`
}

// NewMetadata creates a new Metadata.
func NewMetadata(name, owner string, date time.Time) Metadata {
	return Metadata{
		Name:    name,
		Owner:   owner,
		Created: date.Unix(),
	}
}

// UnpackMetadata reads from r to unpack the metadata of a stream.
func UnpackMetadata(r io.Reader) (Metadata, error) {
	decoder := json.NewDecoder(r)
	var meta Metadata
	if err := decoder.Decode(&meta); err != nil {
		return Metadata{}, err
	}

	if err := Metadata(meta).valid(); err != nil {
		return Metadata{}, err
	}

	return meta, nil
}

// JSON creates a json compatible representation of m.
func (m Metadata) JSON() (string, error) {
	bs, err := json.Marshal(m)
	return string(bs), err
}

func (m Metadata) valid() error {
	if err := ValidName(m.Name); err != nil {
		return err
	}

	if err := ValidOwner(m.Owner); err != nil {
		return err
	}

	return nil
}
