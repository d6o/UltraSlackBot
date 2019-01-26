package logger

import (
	"context"
	"github.com/disiqueira/ultraslackbot/internal/bot"
	usbctx "github.com/disiqueira/ultraslackbot/internal/context"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Execute(ctx context.Context, message slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	usbctx.OutLogger(ctx).Printf("Name: %s Data: %+v\n", message.Name(), message.Data())
	return nil, nil
}

func (l *Logger) Name() string {
	return "Logger"
}
