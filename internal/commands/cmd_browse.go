package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NullEntity2/blogaggregator/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Arguments) > 0 {
		limit, _ = strconv.Atoi(cmd.Arguments[0])
	}

	params := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.Queries.GetPostsByUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error getting feeds with username: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Name: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description)
	}
	return nil
}
