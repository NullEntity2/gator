package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/NullEntity2/blogaggregator/internal/database"
)

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("unfollow requires one argument")
	}
	url := cmd.Arguments[0]

	feed, err := s.Queries.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	delFfParams := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	err = s.Queries.DeleteFeedFollow(context.Background(), delFfParams)
	if err != nil {
		return err
	}

	fmt.Println("Follow deleted")
	return nil
}
