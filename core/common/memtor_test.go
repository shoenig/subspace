// Author hoenig

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MemTorrent(t *testing.T) {
	tor := &MemTorrent{}

	dummy := map[string]int{
		"alpha": 1,
		"beta":  2,
		"gamma": 3,
	}

	encoder := json.NewEncoder(tor)
	err := encoder.Encode(dummy)
	require.NoError(t, err, "failed to encode dummy data")

	decoder := json.NewDecoder(tor)
	unpacked := map[string]int{}
	err = decoder.Decode(&unpacked)
	require.NoError(t, err, "failed to decode dummy data")
	require.Equal(t, 1, unpacked["alpha"])
	require.Equal(t, 2, unpacked["beta"])
	require.Equal(t, 3, unpacked["gamma"])
}
