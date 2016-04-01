// Author hoenig

package main

import (
	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master"
)

func main() {
	filename := config.MustParseFlags()
	config := master.MustLoadConfig(filename)
	server := master.NewMaster(config)
	server.Start()
}
