package main

import (
	"os"

	"github.com/davidovich/summon"

	// change this import to point to your hosted module
	"github.com/davidovich/summon-example-assets/boxer"
)

func main() {
	os.Exit(summon.Main(os.Args, boxer.Box))
}
