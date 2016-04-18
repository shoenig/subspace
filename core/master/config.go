// Author hoenig

package master

import (
	"encoding/json"
	"log"

	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master/state"
)

// MustLoadConfig will load all the options of Config from the file
// specified by filename. Any error encountered will be fatal.
func MustLoadConfig(filename string) *Config {
	var cfg Config
	if err := config.ReadConfig(filename, &cfg); err != nil {
		log.Fatal("could not load config file:", err)
	}
	return &cfg

}

// Config is a representation of the on-disk config file for subspace-master.
type Config struct {
	APIBindAddr string         `json:"api.bind.address"`
	DHTBindAddr string         `json:"dht.bind.address"`
	MasterPeers config.Masters `json:"masters"`
	Raft        state.Config   `json:"raft"`
}

func (c *Config) String() string {
	bs, err := json.MarshalIndent(c, " ", " ")
	if err != nil {
		log.Println("failed to marshal config:", err)
		return "<ERROR>"
	}
	return string(bs)
}
