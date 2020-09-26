package main

import (
	"log"

	"github.com/martinnirtl/dockma/internal/commands"
)

func main() {
	log.SetPrefix("dockma: ")

	commands.Execute()
}
