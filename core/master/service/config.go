// Author hoenig

package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

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

func resolve(address string) (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// RaftMembers returns the members of the raft with resolved addresses and excluding
// "this" instance, which the raft library does not want.
func (c *Config) RaftMembers() ([]string, error) {
	me, err := resolve(c.Raft.BindAddress)
	if err != nil {
		return nil, err
	}

	peers := []string{}
	for _, master := range c.MasterPeers {
		host := fmt.Sprintf("%s:%d", master.Host, master.RaftPort)
		ip, err := resolve(host)
		if err != nil {
			return nil, err
		}
		if ip != me {
			peers = append(peers, ip)
		}
	}
	return peers, nil
}
