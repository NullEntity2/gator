package commands

import (
	"context"

	"github.com/NullEntity2/blogaggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, c Command) error {
		user, err := s.Queries.GetUserByName(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, c, user)
	}
}
