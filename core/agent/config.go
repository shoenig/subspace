// Author hoenig

package agent

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

type Config struct {
	APIBindAddr     string              `json:"api.bind.address"`
	TorrentBindAddr string              `json:"torrent.bind.address"`
	DataDir         string              `json:"data.dir"`
	MasterPeers     []config.MasterPeer `json:"master.peers"`
}

func (c *Config) String() string {
	bs, err := json.MarshalIndent(c, " ", " ")
	if err != nil {
		log.Println("failed to marshal config:", err)
		return "{}"
	}
	return string(bs)
}
