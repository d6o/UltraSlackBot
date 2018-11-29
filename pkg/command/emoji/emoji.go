package emoji

import (
	"fmt"
	"strings"

	"github.com/hackebrot/turtle"
	"github.com/spf13/cobra"
)

const (
	example = `
		# Search emojis related with computers
		!emoji computer`
)

func NewEmojiCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "emoji",
		Short:   "Search for emojis",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := search(strings.Join(args, " "))
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
		},
	}

	return c
}

func search(q string) (string, error) {
	emojis := turtle.Search(q)
	if emojis == nil {
		return "", fmt.Errorf("no emojis found for search: %s", q)
	}

	var msgList []string
	for _, e := range emojis {
		msgList = append(msgList, fmt.Sprintf("%s: %s", e.Name, e.String()))
	}

	return strings.Join(msgList, "\n"), nil
}
