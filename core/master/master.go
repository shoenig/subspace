// Author hoenig

package master

import (
	"fmt"
	"log"
	"time"

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
func (s *Master) Start(bootstrap bool) {
	log.Println("-- subspace-master is starting --")
	log.Println("master config is", s.config)
	log.Println("master will force-start as leader:", bootstrap)

	// -- startup dht --
	dhtServer, err := dht.NewServer(&dht.ServerConfig{
		Addr: s.config.DHTBindAddr,
	})
	if err != nil {
		log.Fatal("failed to start dht server:", err)
	}
	log.Println("dht server is", dhtServer)

	// -- startup raft --
	s.raft, err = state.NewMyRaft(bootstrap, s.config.Raft)
	if err != nil {
		log.Fatal("failed to start raft:", err)
	}

	for range time.Tick(3 * time.Second) {
		log.Println(dhtStats(dhtServer))
	}
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
