package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/NullEntity2/blogaggregator/internal/commands"
	"github.com/NullEntity2/blogaggregator/internal/config"
	"github.com/NullEntity2/blogaggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Read()

	db, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		fatal(err)
	}

	queries := database.New(db)

	s := &commands.State{
		Config:  &cfg,
		Queries: queries,
	}

	cmds := commands.New()

	if len(os.Args) < 2 {
		fatal(errors.New("Error: you must provide an argument"))
	}

	cmd := commands.Command{
		Command:   os.Args[1],
		Arguments: os.Args[2:],
	}
	if err := cmds.Run(s, cmd); err != nil {
		fatal(fmt.Errorf("Error executing command: %w", err))
	}
}

func fatal(e error) {
	fmt.Println(e.Error())
	os.Exit(1)
}
