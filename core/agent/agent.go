// Author hoenig

package agent

import (
	"log"
	"net/http"

	"github.com/anacrolix/torrent"
	"github.com/shoenig/subspace/core/master"
)

// An Agent is the daemon that runs on every machine where torrents may be
// produced or downloaded.
type Agent struct {
	config *Config
	api    *http.Server
}

// NewAgent creates a new Agent.
func NewAgent(config *Config) *Agent {
	return &Agent{
		config: config,
	}
}

// Start starts an Agent.
func (a *Agent) Start() {
	log.Println("-- subspace-agent is starting --")
	log.Println("config is", a.config)

	peerID := GeneratePeerID("") // todo configurable

	tcfg := &torrent.Config{
		DataDir:              a.config.DataDir,
		ListenAddr:           a.config.TorrentBindAddr,
		DisableTrackers:      true,
		DisablePEX:           false,
		NoDHT:                false,
		NoUpload:             false,
		Seed:                 true,
		PeerID:               peerID,
		DisableUTP:           false,
		DisableTCP:           false,
		ConfigDir:            "", // not used
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

	mclient := master.NewClient(a.config.Masters)

	api := apiServer(peerID, a.config.APIBindAddr, a.config.Masters, mclient)

	log.Fatal(api.ListenAndServe())
}
