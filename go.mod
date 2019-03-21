module github.com/davidovich/summon-example-assets

// Note that this go.mod file is typically not needed in an asset repo
// as exemplified by the scaffolder. We use this as an aid
// for summon's development.

require (
	github.com/davidovich/summon v0.1.0
	github.com/gobuffalo/buffalo-plugins v1.13.1 // indirect
	github.com/gobuffalo/packr/v2 v2.0.7
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.3 // indirect
	golang.org/x/sys v0.0.0-20190310054646-10058d7d4faa // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
)

// for dev:
// make sure you have summon checked out as a sibling
replace github.com/davidovich/summon => ../summon

//
// replace github.com/gobuffalo/packr/v2 => ../github.com/gobuffalo/packr/v2
