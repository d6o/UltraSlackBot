package logger

import (
	"context"
	"log"

	"github.com/disiqueira/ultraslackbot/internal/bot"
	usbctx "github.com/disiqueira/ultraslackbot/internal/context"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

type (
	Logger struct {
		logger *log.Logger
	}
)

func New(logger *log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Execute(ctx context.Context, message slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	usbctx.OutLogger(ctx).Printf("Name: %s Data: %+v\n", message.Name(), message.Data())
	return nil, nil
}

func (l *Logger) Name() string {
	return "Logger"
}
