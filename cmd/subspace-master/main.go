// Author hoenig

package main

import (
	"flag"

	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master"
	"github.com/shoenig/subspace/core/master/service"
)

func main() {
	bootstrap := flag.Bool("bootstrap", false, "use bootstrap to force-start as leader")
	filename := config.MustParseFlags()
	config := service.MustLoadConfig(filename)
	server := master.NewMaster(config)
	server.Start(*bootstrap)
}
