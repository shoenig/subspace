// Author hoenig

package main

import (
	"flag"

	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/subspace/core/master"
)

func main() {
	bootstrap := flag.Bool("bootstrap", false, "use bootstrap to force-start as leader")
	filename := config.MustParseFlags()
	config := master.MustLoadConfig(filename)
	server := master.NewMaster(config)
	server.Start(*bootstrap)
}
