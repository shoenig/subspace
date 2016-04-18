// Author hoenig

package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/config"
)

// Torrentify uses the path of a Generation to build the actual torrent
// which can be downloaded by clients.
func Torrentify(
	peerID string,
	masters config.Masters,
	timestamp time.Time,
	gen stream.Generation,
	workers int) (*metainfo.MetaInfo, error) {
	log.Println("[torrent] setting up the builder")
	builder := metainfo.Builder{}

	// 1) add master dht nodes
	builder.AddDhtNodes(masters.DHT())

	log.Println("[torrent] setting metadata")
	// 2) set meta information
	builder.SetName(gen.Stream + "_" + timestamp.Format(time.RFC3339))
	builder.SetComment(gen.Comment)
	builder.SetCreatedBy(peerID)
	builder.SetCreationDate(time.Now().UTC())

	log.Println("[torrent] adding file content")
	// 3) finally add file content
	err := filepath.Walk(gen.Path, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to stat item: %v", err)
		}

		// directories in torrents are implicit
		if !info.IsDir() {
			log.Println("[torrent] adding file path:", fpath)
			builder.AddFile(fpath)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk bundle content: %v", err)
	}

	log.Println("[torrent] executing builder.Submit")
	// 4) create a Batch, which is all preperation for creating a torrent
	batch, err := builder.Submit()
	if err != nil {
		return nil, fmt.Errorf("failed to create batch from bundle: %v", err)
	}

	// 5) start creating the torrent, in memory
	log.Println("[torrent] executing batch.Start")
	mtor := NewMemTorrent()
	errC, progC := batch.Start(mtor, workers)

	log.Println("[torrent] batch.Start has begun, waiting for completion")

	// 6) wait on torrent creation, printing out stats along the way
WAIT:
	for {
		select {
		case err := <-errC:
			if err != nil {
				return nil, fmt.Errorf("failed to create torrent: %v", err)
			}
			break WAIT
		case n := <-progC:
			log.Println("create torrent processed bytes:", n)
		}
	}

	log.Println("[torrent] loading the metainfo from memory")
	// 7) return metainfo on the torrent file we created
	return metainfo.Load(mtor)
}
