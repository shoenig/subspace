// Author hoenig

package state

import (
	"testing"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/stretchr/testify/require"
)

func Test_MyFSM_AddCopyDelete(t *testing.T) {
	fsm := NewMyFSM()

	// add 3 streams
	fsm.AddStreams(
		stream.Stream{
			Info: stream.Info{
				Name:  "stream1",
				Owner: "devops",
			},
		},
		stream.Stream{
			Info: stream.Info{
				Name:  "stream2",
				Owner: "releng",
			},
		},
		stream.Stream{
			Info: stream.Info{
				Name:  "stream3",
				Owner: "delivery",
			},
		},
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
		stream.Stream{Info: stream.Info{Name: "stream3"}},
		stream.Stream{Info: stream.Info{Name: "stream1"}},
	)
	require.Equal(t, 1, len(fsm.streams))
	require.Contains(t, fsm.streams, "stream2")
}
