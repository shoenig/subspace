// Author hoenig

package agent

import (
	"bytes"
	"math/rand"
	"os"
	"time"
)

var hostname = mustGetHostname()
var random = initRandom()

func mustGetHostname() string {
	if h, err := os.Hostname(); err != nil {
		panic(err)
	} else {
		return h
	}
}

// Go uses Seed(1) for the default Rand, we need better randomness
func initRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GeneratePeerID generates an ascii peer id. If prefix is provied
// it is prefixed with that string, otherwise the hostname of the
// machine is used as the prefix.
func GeneratePeerID(prefix string) string {
	var b bytes.Buffer

	if len([]byte(prefix)) > 19 {
		panic("length of peerID must be fewer than 19 bytes")
	}

	if prefix == "" {
		b.WriteString(hostname)
	} else {
		b.WriteString(prefix)
	}

	b.WriteString("-")

	for i := b.Len(); i < 20; i++ {
		b.WriteByte(byte(random.Int31n(26) + 97))
	}
	return b.String()
}
