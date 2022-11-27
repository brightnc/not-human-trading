package main

import (
	"fmt"

	cmd "github.com/brightnc/not-human-trading/cmd"
	"github.com/brightnc/not-human-trading/protocol"
)

func main() {
	fmt.Printf(`
    ___       ______    _____
   / _  \    |   _  \  |_   _|
  /  __  \   |   __ /   _ | _
 /__/  \__\  |__|      |_____|

   github.com/brightnc/not-human-trading %s, built with Go %s
 `, cmd.Version, cmd.GoVersion)
	// cmd.Execute()
	protocol.ServeREST()
}
