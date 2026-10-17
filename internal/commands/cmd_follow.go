package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NullEntity2/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("follow requires one argument")
	}
	url := cmd.Arguments[0]

	feed, err := s.Queries.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	ff, err := s.Queries.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("Feed Name: %s\n", ff.FeedName)
	fmt.Printf("Username: %s\n", ff.UserName)
	return nil
}
