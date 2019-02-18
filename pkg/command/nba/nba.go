package nba

import (
	"fmt"
	"github.com/disiqueira/ultraslackbot/pkg/nba"
	"github.com/spf13/cobra"
)

type (
	nbaCommand struct {
		client *nba.NBA
	}
)

const (
	template = "%s (%s - %s) x %s (%s - %s)\n%s\n%s\n"
)

func NewNBACommand() *cobra.Command {
	client := nba.New()
	n := nbaCommand{
		client: client,
	}
	c := &cobra.Command{
		Use:   "nba",
		Short: "NBA Games today",
		Args:  cobra.NoArgs,
		Run:   n.Games,
	}

	c.AddCommand(newTeamCommand(client))

	return c
}

func (n *nbaCommand) Games(cmd *cobra.Command, args []string) {
	games, err := n.client.Today()
	if err != nil {
		_, _ = cmd.OutOrStdout().Write([]byte(err.Error()))
		return
	}

	teams, err := n.client.Teams()
	if err != nil {
		_, _ = cmd.OutOrStdout().Write([]byte(err.Error()))
		return
	}

	teamList := map[string]nba.Team{}
	for _, team := range teams.League.Standard {
		teamList[team.TeamID] = team
	}

	msg := "```"
	for _, game := range games.Games {
		msg += fmt.Sprintf(template, teamList[game.HTeam.TeamID].FullName, game.HTeam.Win, game.HTeam.Loss,
			teamList[game.VTeam.TeamID].FullName, game.VTeam.Win, game.VTeam.Loss, game.StartTimeEastern,
			game.Nugget.Text)
	}

	msg += "```"
	_, _ = cmd.OutOrStdout().Write([]byte(msg))
}
