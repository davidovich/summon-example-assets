module github.com/davidovich/summon-example-assets

// Note that this go.mod file is typically not needed in an asset repo
// as exemplified by the scaffolder. We use this as an aid
// for summon's development.

replace github.com/davidovich/summon-example-assets/summon v0.0.0 => ./summon

require (
	github.com/davidovich/summon v0.3.0 // indirect
	github.com/gobuffalo/packr/v2 v2.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c // indirect
	golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect
)
