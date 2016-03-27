// Author hoenig

package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type MasterPeer struct {
	Host    string `json:"host"`
	APIPort int    `json:"api.port"`
	DHTPort int    `json:"dht.port"`
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
