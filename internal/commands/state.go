package commands

import (
	"github.com/NullEntity2/blogaggregator/internal/config"
	"github.com/NullEntity2/blogaggregator/internal/database"
)

type State struct {
	Config  *config.Config
	Queries *database.Queries
}
