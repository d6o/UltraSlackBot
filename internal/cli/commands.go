package cli

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/disiqueira/ultraslackbot/internal/app"
)

const (
	errorExitCode   = 1
	runLongDescription = `
Creates a new connection to Slack and starts handling all the new messages
that receives from it.

It will load all the configuration from the environment vars, validate all
the plugins in the plugin folder, load all the plugins and start the message
handler.`
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "ultraslackbot",
		Short: "UltraSlackBot is your extendable Slack Bot",
	}

	a := app.New()

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Start the Ultra Slack Bot",
		Long:  runLongDescription,
		Run: a.Run,
	}

	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	}
}
