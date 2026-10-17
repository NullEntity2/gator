package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/NullEntity2/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("register requires exactly one argument")
	}
	name := cmd.Arguments[0]

	// check if user already exists
	foundUser, err := s.Queries.GetUserByName(context.Background(), name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("error getting user by name: %w", err)
		}
	}

	if foundUser.ID != uuid.Nil {
		return fmt.Errorf("A user already exists with the name %s", name)
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	createdUser, err := s.Queries.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	s.Config.SetUser(name)

	fmt.Printf("User %s has been created\n", createdUser.Name)
	fmt.Printf("ID: %v\n", createdUser.ID)
	fmt.Printf("CreatedAt: %v\n", createdUser.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", createdUser.UpdatedAt)
	fmt.Printf("Name: %v\n", createdUser.Name)
	return nil
}
