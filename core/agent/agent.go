// Author hoenig

package agent

import (
	"log"
	"net/http"

	"github.com/anacrolix/torrent"
)

type Agent struct {
	config *Config
	api    *http.Server
}

func NewAgent(config *Config) *Agent {
	return &Agent{
		config: config,
	}
}

func (a *Agent) Start() {
	log.Println("-- subspace-agent is starting --")
	log.Println("config is", a.config)

	tcfg := &torrent.Config{
		DataDir:              a.config.DataDir,
		ListenAddr:           a.config.TorrentBindAddr,
		DisableTrackers:      true,
		DisablePEX:           false,
		NoDHT:                false,
		NoUpload:             false,
		Seed:                 true,
		PeerID:               GeneratePeerID(""),
		DisableUTP:           false,
		DisableTCP:           false,
		NoDefaultBlocklist:   true,
		ConfigDir:            "", // eh
		DisableMetainfoCache: true,
		DisableEncryption:    true, // todo false
		DisableIPv6:          true, // todo false
		IPBlocklist:          nil,  // todo configure
		Debug:                true, // todo configure
	}

	tclient, err := torrent.NewClient(tcfg)
	if err != nil {
		log.Fatalln("failed to start torrent client:", err)
	}

	log.Println("i am agent", tclient.PeerID())

	api := apiServer(a.config.APIBindAddr, tclient)
	log.Fatal(api.ListenAndServe())
}
