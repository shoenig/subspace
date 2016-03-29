// Author hoenig

package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type Masters []MasterPeer

func (m Masters) DHT() []string {
	peers := []string{}
	for _, master := range m {
		peers = append(peers, master.DHT())
	}
	return peers
}

type MasterPeer struct {
	Host    string `json:"host"`
	APIPort int    `json:"api.port"`
	DHTPort int    `json:"dht.port"`
}

func (p MasterPeer) DHT() string {
	return fmt.Sprintf("%s:%d", p.Host, p.DHTPort)
}

func (p MasterPeer) API() string {
	return fmt.Sprintf("%s:%d", p.Host, p.APIPort)
}

func MustParseFlags() string {
	configFile := flag.String("config", "", "the configuration json file (required)")
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		log.Fatal("--config was not specified")
	}
	return *configFile
}

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
