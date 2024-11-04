package main

import (
	"log"
	"os"

	"github.com/tanun45/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	st := state{
		cfg: &cfg,
	}

	cmds := commands{
		commandsMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: Not enough arguments")
	}

	commandName := args[1]
	commandArgs := args[2:]
	err = cmds.run(&st, command{
		Name: commandName,
		Args: commandArgs,
	})
	if err != nil {
		log.Fatalf("Error running the command: %v\n", err)
	}
}
