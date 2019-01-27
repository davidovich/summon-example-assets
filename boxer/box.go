package boxer

import (
	"github.com/gobuffalo/packr/v2"
)

// Box captures the files of the assets tree
var Box = packr.New("Summon Box", "../assets")
