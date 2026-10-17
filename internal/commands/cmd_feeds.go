package commands

import (
	"context"
	"fmt"
)

func handlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Queries.GetFeedsWithUsername(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds with username: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Created By: %s\n", feed.UserName)
	}
	return nil
}
