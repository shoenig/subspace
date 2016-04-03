// Author hoenig

package master

import (
	"fmt"
	"log"

	"github.com/anacrolix/torrent/dht"
	"github.com/shoenig/subspace/core/master/state"
)

// Master does things.
type Master struct {
	config *Config
	raft   *state.MyRaft
}

// NewMaster creates a Master which does things.
func NewMaster(config *Config) *Master {
	return &Master{
		config: config,
	}
}

// Start causes the Master to do things.
func (m *Master) Start(bootstrap bool) {
	log.Println("-- subspace-master is starting --")
	log.Println("master config is", m.config)
	log.Println("master will force-start as leader:", bootstrap)

	// -- startup dht --
	dhtServer, err := dht.NewServer(&dht.ServerConfig{
		Addr: m.config.DHTBindAddr,
	})
	if err != nil {
		log.Fatal("failed to start dht server:", err)
	}
	log.Println("dht server is", dhtServer)

	// -- startup raft --
	m.raft, err = state.NewMyRaft(bootstrap, m.config.Raft)
	if err != nil {
		log.Fatal("failed to start raft:", err)
	}

	api := apiServer(m.config.APIBindAddr)
	log.Fatal(api.ListenAndServe())

}

func dhtStats(dht *dht.Server) string {
	stats := dht.Stats()
	return fmt.Sprintf(
		"nodes: %d, good nodes: %d bad nodes: %d, txs: %d, confirms: %d",
		stats.Nodes,
		stats.GoodNodes,
		stats.BadNodes,
		stats.OutstandingTransactions,
		stats.ConfirmedAnnounces,
	)
}
