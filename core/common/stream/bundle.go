// Author hoenig

package stream

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/anacrolix/torrent/metainfo"
)

// A Pack is a Bundle + MagnetURI for its contents, after a torrent
// of the content has been created and is ready for seeding.
type Pack struct {
	Bundle
	MagnetURI string `json:"magnet"`
}

// NewPack creates a new Pack.
func NewPack(b Bundle, magnet string) Pack {
	return Pack{
		Bundle:    b,
		MagnetURI: magnet,
	}
}

// UnpackPack unpacks a json representation of a Pack.
func UnpackPack(r io.Reader) (Pack, error) {
	decoder := json.NewDecoder(r)
	var pack Pack
	if err := decoder.Decode(&pack); err != nil {
		return Pack{}, err
	}
	if err := pack.Bundle.valid(); err != nil {
		return Pack{}, err
	}
	// verify the magnet uri is at least parsable
	if _, err := metainfo.ParseMagnetURI(pack.MagnetURI); err != nil {
		return Pack{}, err
	}

	return pack, nil
}

// JSON returns the json representation of a Pack.
func (p Pack) JSON() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// Bundle represents the publication of a new generation of content on a Stream.
type Bundle struct {
	Info
	Path    string `json:"path"`
	Comment string `json:"comment"`
}

// UnpackBundle unpacks a json representation of a Bundle.
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

func (b Bundle) valid() error {
	if err := b.Info.valid(); err != nil {
		return err
	}
	if b.Path == "" {
		return fmt.Errorf("bundle cannot have empty path")
	}
	return nil
}
