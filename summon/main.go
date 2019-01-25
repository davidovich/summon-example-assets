package main

import (
	"os"

	"github.com/davidovich/summon"
)

/*
Dev runs the app with the assets he wants

config allows to introduce aliases

filesystem can be accessed directly
*/

func main() {
	os.Exit(summon.Main(os.Args))
}
