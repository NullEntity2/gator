package commands

import "fmt"

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	v, ok := c.Commands[cmd.Command]
	if !ok {
		return fmt.Errorf("could not find command with name \"%s\"", cmd.Command)
	}
	return v(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	_, ok := c.Commands[name]
	if ok {
		return fmt.Errorf("a command already exists with the name %s", name)
	}
	c.Commands[name] = f
	return nil
}

func New() *Commands {
	cmds := Commands{
		Commands: map[string]func(*State, Command) error{},
	}
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)
	cmds.Register("reset", handlerReset)
	cmds.Register("users", handlerUsers)
	cmds.Register("agg", handlerAgg)
	cmds.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.Register("feeds", handlerFeeds)
	cmds.Register("follow", middlewareLoggedIn(handlerFollow))
	cmds.Register("following", middlewareLoggedIn(handlerFollowing))
	cmds.Register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.Register("browse", middlewareLoggedIn(handlerBrowse))
	return &cmds
}
