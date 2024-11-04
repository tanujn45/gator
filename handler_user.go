package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login handler expects a single argument, the username!")
	}

	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' has been set\n", username)
	return nil
}
