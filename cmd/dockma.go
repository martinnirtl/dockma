package main

import (
	"github.com/martinnirtl/dockma/internal/commands"
)

// TODO set version, commit and date with build flags and in dev mode
// var (
// 	version = "dev"
// 	commit  = "none"
// 	date    = "unknown"
// )

func main() {
	commands.Execute()
}
