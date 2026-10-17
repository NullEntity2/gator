package commands

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	// check if user already exists
	err := s.Queries.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to delete users: %w", err)
	}

	fmt.Printf("Database has been reset\n")
	return nil
}
