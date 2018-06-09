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
	"github.com/disiqueira/ultraslackbot/internal/conf"
	"github.com/disiqueira/ultraslackbot/pkg/handlers/logger"
	"github.com/disiqueira/ultraslackbot/internal/handlers/admin"
)

type (
	App struct {}
)

const (
	successExitCode = 0
	errorExitCode   = 1
	slackTokenEnvVar   = "SLACKTOKEN"
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

	specs := conf.Load()

	outLogger.Printf("Configurations: %+v", specs.All())

	slackToken, ok := specs.Get(slackTokenEnvVar)
	if !ok {
		errLogger.Printf("config missing %s", slackTokenEnvVar)
		os.Exit(errorExitCode)
	}

	pluginReader := plugin.New()
	loadedHandlers, err := pluginReader.Load(errLogger, specs)
	if err != nil {
		errLogger.Printf("Loading plugins error: %s", err.Error())
	}

	slackClient := slack.New(slackToken.(string))
	b := bot.New(slackClient)
	b.SetHandlers(append(a.defaultHandlers(ctx, b), loadedHandlers...))
	b.Run(ctx)
	os.Exit(successExitCode)
}

func (a *App) defaultHandlers(ctx context.Context, b *bot.Bot) []bot.Handler {
	return []bot.Handler{
		logger.New(usbCtx.OutLogger(ctx)),
		admin.New(b),
		}
}
