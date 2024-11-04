package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	commandFunc, ok := c.commandsMap[cmd.Name]
	if !ok {
		return errors.New("Command does not exist")
	}
	return commandFunc(s, cmd)
}
