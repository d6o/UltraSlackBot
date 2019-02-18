package nba

import (
	"fmt"
	"github.com/disiqueira/ultraslackbot/pkg/nba"
	"github.com/spf13/cobra"
)

type (
	nbaTeamCommand struct {
		client *nba.NBA
	}
)

var (
	templateTeam = "```Name: %s\n" +
		"ID: %s\n" +
		"Wins: %s\n" +
		"Losses: %s\n" +
		"Conference Rank: %s\n" +
		"Division Rank: %s\n" +
		"Win Percentage: %s\n" +
		"Loss Percentage: %s\n" +
		"Games Behind: %s\n" +
		"Division Games Behind: %s\n" +
		"Conference: %s - %s\n" +
		"Division: %s - %s\n" +
		"Home: %s - %s\n" +
		"Away: %s - %s\n" +
		"Last 10: %s - %s\n" +
		"Streak: %s %s\n" +
		"Tie Breaker Points: %s ```"
)

func newTeamCommand(client *nba.NBA) *cobra.Command {
	n := nbaTeamCommand{
		client: client,
	}
	c := &cobra.Command{
		Use:   "team",
		Short: "Infos about a NBA team",
		Args:  cobra.MinimumNArgs(1),
		Run:   n.Team,
	}

	return c
}

func (c *nbaTeamCommand) Team(cmd *cobra.Command, args []string) {
	standings, err := c.client.Standing()
	if err != nil {
		_, _ = cmd.OutOrStdout().Write([]byte(err.Error()))
		return
	}

	teams, err := c.client.Teams()
	if err != nil {
		_, _ = cmd.OutOrStdout().Write([]byte(err.Error()))
		return
	}

	teamList := map[string]nba.Team{}
	for _, team := range teams.League.Standard {
		teamList[team.TeamID] = team
	}

	for _, team := range standings.League.Standard.Teams {
		if team.TeamID == args[0] {
			streak := "losses"
			if team.IsWinStreak {
				streak = "wins"
			}
			msg := fmt.Sprintf(templateTeam, teamList[team.TeamID].FullName, team.TeamID, team.Win, team.Loss, team.ConfRank,
				team.DivRank, team.WinPctV2, team.LossPctV2, team.GamesBehind, team.DivGamesBehind, team.ConfWin,
				team.ConfLoss, team.DivWin, team.DivLoss, team.HomeWin, team.HomeLoss, team.AwayWin, team.AwayLoss,
				team.LastTenWin, team.LastTenLoss, team.Streak, streak, team.TieBreakerPts)

			_, _ = cmd.OutOrStdout().Write([]byte(msg))
			return
		}
	}
}
