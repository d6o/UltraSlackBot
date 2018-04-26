package app

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	usbCtx "github.com/disiqueira/ultraslackbot/internal/context"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/internal/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/handlers"
)

type (
	App struct {
	}
)

const (
	successExitCode = 0
	errorExitCode   = 1
	slackTokenEnvVar   = "SLACK_TOKEN"
)

func New() *App {
	return &App{}
}

func (a *App) Run(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	outLogger := log.New(os.Stdout, "", 0)
	errLogger := log.New(os.Stderr, "", 0)

	ctx = usbCtx.WithOutLogger(ctx, outLogger)
	ctx = usbCtx.WithErrLogger(ctx, errLogger)

	slackToken := os.Getenv(slackTokenEnvVar)
	if slackToken == "" {
		errLogger.Printf("you need to inform your slack token. export %s={YOUR_TOKEN}", slackTokenEnvVar)
		os.Exit(errorExitCode)
	}

	pluginReader := plugin.New()
	loadedHandlers, err := pluginReader.Load(errLogger)
	if err != nil {
		errLogger.Printf("Loading plugins error: %s", err.Error())
	}

	slackClient := slack.New(slackToken)
	b := bot.New(slackClient, append(a.defaultHandlers(ctx), loadedHandlers...))
	b.Run(ctx)
	os.Exit(successExitCode)
}

func (a *App) defaultHandlers(ctx context.Context) []bot.Handler {
	return []bot.Handler{
		handlers.NewLogger(usbCtx.OutLogger(ctx)),
		}
}
