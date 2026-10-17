package commands

import (
	"context"
	"fmt"

	"github.com/NullEntity2/blogaggregator/internal/database"
)

func handlerFollowing(s *State, cmd Command, user database.User) error {
	ffs, err := s.Queries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %w", err)
	}

	for _, ff := range ffs {
		fmt.Printf("Feed Name: %s\n", ff.FeedName)
	}

	return nil
}
