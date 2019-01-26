package main

import (
	"os"

	"github.com/davidovich/summon"
	"github.com/gobuffalo/packr/v2"
)

/*
Dev runs the app with the assets he wants

config allows to introduce aliases

filesystem can be accessed directly
*/

func main() {
	box := packr.New("Box", "./assets")
	os.Exit(summon.Main(os.Args, box))
}
