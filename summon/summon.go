package main

import (
	"os"
	"runtime/debug"

	"github.com/davidovich/summon"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	// Box captures the files of the assets tree
	box := packr.New("Summon Box", "../assets")
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("")
	}
	println("mod", bi.Main.Path, bi.Main.Version)
	for _, d := range bi.Deps {
		println("dep", d.Path, d.Version, d.Sum)
		if r := d.Replace; r != nil {
			println("=>", r.Path, r.Version, r.Sum)
		}
	}
	os.Exit(summon.Main(os.Args, box))
}
