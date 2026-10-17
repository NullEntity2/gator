package commands

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("login requires exactly one argument")
	}
	name := cmd.Arguments[0]

	_, err := s.Queries.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("error getting users by name: %w", err)
	}

	if err := s.Config.SetUser(name); err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Printf("User has been set to %s\n", cmd.Arguments[0])
	return nil
}
