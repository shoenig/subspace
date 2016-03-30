// Author hoenig

package master

import (
	"encoding/json"
	"log"

	"github.com/shoenig/subspace/core/config"
)

func MustLoadConfig(filename string) *Config {
	var cfg Config
	if err := config.ReadConfig(filename, &cfg); err != nil {
		log.Fatal("could not load config file:", err)
	}
	return &cfg

}

// Config is a representation of the on-disk config file for subspace-master.
type Config struct {
	// torrent
	APIBindAddr string         `json:"api.bind.address"`
	DHTBindAddr string         `json:"dht.bind.address"`
	MasterPeers config.Masters `json:"master.peers"`

	// raft
	SingleMasterMode bool `json:"single.master.mode"`
}

func (c *Config) String() string {
	bs, err := json.MarshalIndent(c, " ", " ")
	if err != nil {
		log.Println("failed to marshal config:", err)
		return "{}"
	}
	return string(bs)
}
