package commands

import (
	"context"
	"fmt"
)

func handlerUsers(s *State, cmd Command) error {
	// check if user already exists
	users, err := s.Queries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
