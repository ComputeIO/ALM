// Package main stands for default entrance
package main

import (
	_ "embed"
	"github.com/ComputeIO/alm/pkg/config"
)

// Version semantic version string
var Version string = "dev"

func main() {
	config.Get()
}
