// Author hoenig

package master

import (
	"fmt"
	"log"
	"time"

	"github.com/anacrolix/torrent/dht"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {
	log.Println("-- subspace-master is starting --")
	log.Println("master config is", s.config)

	dhtServer, err := dht.NewServer(&dht.ServerConfig{
		Addr: s.config.DHTBindAddr,
	})

	if err != nil {
		log.Fatal("failed to start dht server:", err)
	}

	log.Println("dht server is", dhtServer)

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
