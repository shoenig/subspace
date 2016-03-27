// Author hoenig

package main

import (
	"github.com/shoenig/subspace/core/agent"
	"github.com/shoenig/subspace/core/config"
)

func main() {
	filename := config.MustParseFlags()
	config := agent.MustLoadConfig(filename)
	agent := agent.NewAgent(config)
	agent.Start()
}
