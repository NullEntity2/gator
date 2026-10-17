package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NullEntity2/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return errors.New("addfeed requires two arguments")
	}
	name := cmd.Arguments[0]
	url := cmd.Arguments[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.Queries.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error creating feed: %w", err)
	}

	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.Queries.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}

	fmt.Printf("%v\n", feed)
	return nil
}
