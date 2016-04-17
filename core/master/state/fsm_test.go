// Author hoenig

package state

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/hashicorp/raft"
	"github.com/shoenig/subspace/core/common/stream"
	"github.com/stretchr/testify/require"
)

func Test_MyFSM_AddCopyDelete(t *testing.T) {
	fsm := NewMyFSM()

	// add 3 streams
	fsm.AddStreams(
		stream.NewStream("stream1", "devops"),
		stream.NewStream("stream2", "releng"),
		stream.NewStream("stream3", "delivery"),
	)

	require.Equal(t, 3, len(fsm.streams))
	require.Contains(t, fsm.streams, "stream1")
	require.Contains(t, fsm.streams, "stream2")
	require.Contains(t, fsm.streams, "stream3")

	// copy the 3 streams
	copied := fsm.CopyStreams()
	require.Equal(t, 3, len(copied))
	require.Contains(t, fsm.streams, copied[0].Name)
	require.Equal(t, fsm.streams[copied[0].Name].Owner, copied[0].Owner)
	require.Contains(t, fsm.streams, copied[1].Name)
	require.Equal(t, fsm.streams[copied[1].Name].Owner, copied[1].Owner)
	require.Contains(t, fsm.streams, copied[2].Name)
	require.Equal(t, fsm.streams[copied[2].Name].Owner, copied[2].Owner)

	// delete 2 of the 3 streams
	fsm.DeleteStreams(
		stream.NewStream("stream3", "delivery"),
		stream.NewStream("stream1", "devops"),
	)
	require.Equal(t, 1, len(fsm.streams))
	require.Contains(t, fsm.streams, "stream2")
}

func Test_MyFSM_Do_Streams(t *testing.T) {
	fsm := NewMyFSM()

	action1 := Action{
		Command: AddStreams,
		Streams: []stream.Stream{
			stream.NewStream("stream1", "devops"),
			stream.NewStream("stream2", "sem"),
			stream.NewStream("stream3", "squall"),
			stream.NewStream("stream4", "ops"),
		},
	}

	fsm.Do(action1)
	require.Equal(t, 4, len(fsm.streams))
	require.Equal(t, "devops", fsm.streams["stream1"].Owner)
	require.Equal(t, "sem", fsm.streams["stream2"].Owner)
	require.Equal(t, "squall", fsm.streams["stream3"].Owner)
	require.Equal(t, "ops", fsm.streams["stream4"].Owner)

	action2 := Action{
		Command: DeleteStreams,
		Streams: []stream.Stream{
			stream.NewStream("stream1", ""),
			stream.NewStream("stream4", ""),
		},
	}

	fsm.Do(action2)
	require.Equal(t, 2, len(fsm.streams))
	require.Equal(t, "sem", fsm.streams["stream2"].Owner)
	require.Equal(t, "squall", fsm.streams["stream3"].Owner)
}

// for mocking out an io.ReadCloser from which a snapshot can be read
type byteReadCloser struct {
	buf *bytes.Buffer
}

func (brc byteReadCloser) Read(p []byte) (int, error) {
	return brc.buf.Read(p)
}

func (brc byteReadCloser) Close() error {
	return nil
}

func Test_MyFSM_ApplySnapshot_Streams(t *testing.T) {
	fsm := NewMyFSM()

	data, err := json.Marshal(Action{
		Command: AddStreams,
		Streams: []stream.Stream{stream.NewStream("stream1", "devops")},
	})
	require.NoError(t, err, "json marshal log failed")

	log := &raft.Log{
		Data: data,
	}

	// -- apply --
	fsm.Apply(log) // did not panic

	// -- snapshot --
	snap, err := fsm.Snapshot()
	require.NoError(t, err, "snapshot failed")
	require.NotNil(t, snap, "snapshot should not be nil")
}

func Test_MyFSM_Restore_Streams(t *testing.T) {
	// create a fresh fsm with no data
	fsm := NewMyFSM()

	data, err := json.Marshal([]stream.Stream{stream.NewStream("stream1", "devops")})
	require.NoError(t, err, "json marshal log failed")

	snap := byteReadCloser{
		buf: bytes.NewBuffer(data),
	}

	err = fsm.Restore(snap)
	require.NoError(t, err, "restore from snapshot failed")
}
