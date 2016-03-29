// Author hoenig

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Masters_DHT(t *testing.T) {
	masters := Masters([]MasterPeer{
		{
			Host:    "1.2.3.4",
			APIPort: 1555,
			DHTPort: 1666,
		},
		{
			Host:    "2.3.4.5",
			APIPort: 7888,
			DHTPort: 8999,
		},
	})
	dht := masters.DHT()
	require.Equal(t, 2, len(dht))
	require.Contains(t, dht, "1.2.3.4:1666")
	require.Contains(t, dht, "2.3.4.5:8999")
}
