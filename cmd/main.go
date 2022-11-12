package main

import (
	"fmt"

	"github.com/brightnc/not-human-trading/cmd/cmds"
)

func main() {
	fmt.Printf(`
    ___       ______    _____
   / _  \    |   _  \  |_   _|
  /  __  \   |   __ /   _ | _
 /__/  \__\  |__|      |_____|

   github.com/brightnc/not-human-trading %s, built with Go %s
 `, cmds.Version, cmds.GoVersion)
	cmds.Execute()
}
