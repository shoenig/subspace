// Author hoenig

package agent

import (
	"encoding/json"
	"log"

	"github.com/shoenig/subspace/core/config"
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

// Config represents the available runtime options.
type Config struct {
	Masters         config.Masters `json:"masters"`
	APIBindAddr     string         `json:"api.bind.address"`
	TorrentBindAddr string         `json:"torrent.bind.address"`
	DataDir         string         `json:"data.dir"`
}

func (c *Config) String() string {
	bs, err := json.MarshalIndent(c, " ", " ")
	if err != nil {
		log.Println("failed to marshal config:", err)
		return "{}"
	}
	return string(bs)
}
