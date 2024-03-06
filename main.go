package main

import (
	_ "embed"

	"github.com/Infinity-Green/inf/command/root"
	"github.com/Infinity-Green/inf/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
