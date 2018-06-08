package handlers

import (
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"log"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

type (
	Logger struct {
		logger *log.Logger
	}
)

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		logger:logger,
	}
}

func (l *Logger) Start(spec bot.Specs) error {
	return nil
}

func (l *Logger) Execute(message slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	log.Printf("Name: %s Data: %+v\n",message.Name(), message.Data())
	return nil, nil
}

func (l *Logger) Name() string {
	return "Logger"
}
