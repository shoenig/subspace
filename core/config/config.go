// Author hoenig

package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

// Masters is a list of connected, all-knowing MasterPeers.
type Masters []MasterPeer

// DHT returns the list of DHT addresses of the static list of MasterPeers.
func (m Masters) DHT() []string {
	peers := []string{}
	for _, master := range m {
		peers = append(peers, master.DHT())
	}
	return peers
}

// A MasterPeer is a static, all-knowing entity that all clients know about so that they
// can join the swarm and make API requests.
type MasterPeer struct {
	Host          string `json:"host"`
	APIPort       int    `json:"api.port"`
	APIDisableTLS bool   `json:"api.disable.tls"`
	DHTPort       int    `json:"dht.port"`
	RaftPort      int    `json:"raft.port"`
}

// DHT returns the address for torrent and swarm activity of p.
func (p MasterPeer) DHT() string {
	return fmt.Sprintf("%s:%d", p.Host, p.DHTPort)
}

// API returns the address for clients to submit api requests.
func (p MasterPeer) API(endpoint string) string {
	protocol := "https"
	if p.APIDisableTLS {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s:%d/%s", protocol, p.Host, p.APIPort, endpoint)
}

// MustParseFlags parses the command line flags and returns the path to a config file.
func MustParseFlags() string {
	configFile := flag.String("config", "", "the configuration json file (required)")
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		log.Fatal("--config was not specified")
	}
	return *configFile
}

// ReadConfig reads configuration out of the named file into c.
func ReadConfig(filename string, c interface{}) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bs, c); err != nil {
		return err
	}
	return nil
}
