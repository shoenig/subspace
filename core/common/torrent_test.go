// Author hoenig

package common

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/config"
	"github.com/stretchr/testify/require"
)

func tempFile(t *testing.T, content string) *os.File {
	tempfile, err := ioutil.TempFile("", "state-")
	require.NoError(t, err, "could not create tempfile")
	err = tempfile.Sync()
	require.NoError(t, err, "failed to fsync tempfile")
	err = tempfile.Close()
	require.NoError(t, err, "failed to close tempfile")
	return tempfile
}

func Test_Torrentify(t *testing.T) {
	f := tempFile(t, "this is the content of a file!")
	defer func(t *testing.T, file *os.File) {
		require.NoError(t, os.Remove(file.Name()), "failed to remove tempfile")
	}(t, f)

	masters := []config.MasterPeer{
		{Host: "1.2.3.4", DHTPort: 1234},
		{Host: "2.3.4.5", DHTPort: 2345},
	}
	now := time.Date(2016, 04, 17, 22, 03, 0, 0, time.UTC)
	generation := stream.Generation{
		Stream:  "stream1",
		Path:    f.Name(),
		Comment: "Test_Torrentify torrent",
	}

	minfo, err := Torrentify("myPeerID", masters, now, generation, 3)

	require.NoError(t, err, "failed to torrentfy tempfile")
	require.Equal(t, "myPeerID", minfo.CreatedBy)
	require.Equal(t, "Test_Torrentify torrent", minfo.Comment)
	require.Equal(t, int64(256*1024), minfo.Info.PieceLength)
	require.Contains(t, minfo.Nodes, metainfo.Node("1.2.3.4:1234"))
	require.Contains(t, minfo.Nodes, metainfo.Node("2.3.4.5:2345"))
}
