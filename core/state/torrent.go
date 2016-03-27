// Author hoenig

package state

import (
	"github.com/anacrolix/torrent/metainfo"
)

func CreateBundle(path string) (*metainfo.MetaInfo, error) {
	minfo, err := metainfo.LoadFromFile(path)

	return minfo, err
}
