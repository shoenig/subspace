// Author hoenig

package state

import "bytes"

// An in-memory torrent (todo, add compression)
type MemTorrent struct {
	data bytes.Buffer
}

func NewMemTorrent() *MemTorrent {
	return &MemTorrent{}
}

func (t *MemTorrent) Read(p []byte) (n int, err error) {
	return t.data.Read(p)
}

func (t *MemTorrent) Write(p []byte) (n int, err error) {
	return t.data.Write(p)
}
