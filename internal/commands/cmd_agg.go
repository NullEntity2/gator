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

func handlerAgg(s *State, cmd Command) error {
	interval, err := getInterval(cmd)
	if err != nil {
		return fmt.Errorf("failed to getInterval: %w", err)
	}

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error scraping feed: %v\n", err)
		}
	}
}

func scrapeFeeds(s *State) error {
	nextFeed, err := s.Queries.GetNextFeedToFetch(context.Background())
	if err := s.Queries.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("failed to MarkFeedFetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("%s\n", nextFeed.Name)
	for _, v := range feed.Channel.Item {
		fmt.Printf("- %v\n", v.Title)

		parsedTime, err := parseToNullTime(v.PubDate)
		if err != nil {
			fmt.Printf("warning parsing date \"%s\". Using null date.: %v", v.PubDate, err)
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       v.Title,
			Url:         v.Link,
			Description: v.Description,
			PublishedAt: parsedTime,
			FeedID:      nextFeed.ID,
		}
		if _, err := s.Queries.CreatePost(context.Background(), params); err != nil {
			fmt.Printf("warning creating post: %v\n", err)
		}
		fmt.Printf("Post \"%s\" created\n", v.Title)
	}

	return nil
}

func getInterval(cmd Command) (time.Duration, error) {
	if len(cmd.Arguments) != 1 {
		return 0, errors.New("agg requires exactly one argument")
	}
	intervalStr := cmd.Arguments[0]

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		return 0, fmt.Errorf("Could not format as a date: %v", err)
	}

	return interval, nil
}

func parseToNullTime(dateStr string) (sql.NullTime, error) {
	parsedTime, err := time.Parse(time.RFC1123Z, dateStr)
	if err != nil {
		return sql.NullTime{}, fmt.Errorf("error parsing date: %w", err)
	}

	return sql.NullTime{
		Time:  parsedTime,
		Valid: true,
	}, nil
}
