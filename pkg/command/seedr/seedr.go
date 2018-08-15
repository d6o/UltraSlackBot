package seedr

import (
			"github.com/spf13/cobra"
				"github.com/disiqueira/ultraslackbot/pkg/seedr"
)

const (
	example = `
		# Ask a question
		!seedr Who is the president of Brazil?

		# Ask another question
		!seedr What is the distance between the Earth and the Moon?

		# Solve math problems
		!calc 2+2`
)

type (
	seedrManager struct {
		client *seedr.Seedr
	}
)

func NewSeedrCommand(username, password string) *cobra.Command {
	s := seedr.New(username, password)

	c := &cobra.Command{
		Use:     "seedr",
		Short:   "Ask a question to Seedr|Alpha",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := s.Folders()
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"wa", "calc", "ask"},
	}

	return c
}

func newSeedrManager() *seedrManager {
	return &seedrManager{
		client:
	}
}