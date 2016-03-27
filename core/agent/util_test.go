// Author hoenig

package agent

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GeneratePeerID_default(t *testing.T) {
	peerID := GeneratePeerID("")
	require.Equal(t, 20, len([]byte(peerID)), "default peerID was not 20 bytes in length")
}

func Test_GeneratePeerID_custom(t *testing.T) {
	peerID := GeneratePeerID("abc123")
	require.Equal(t, 20, len([]byte(peerID)), "custom peerID was not 20 bytes in length")
	require.True(t, strings.HasPrefix(peerID, "abc123-"), "peerID did not start with prefix")
}

func Test_GeneratePeerID_long(t *testing.T) {
	require.Panics(t, func() { GeneratePeerID("thispeeridiswaytoolongtobeused") })
}
