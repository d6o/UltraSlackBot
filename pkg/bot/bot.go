package bot

import (
	"context"
	usbCtx "github.com/disiqueira/ultraslackbot/internal/context"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"time"
	"errors"
)

type (
	Bot struct {
		token    string
		handlers []Handler
		slack    SlackClient
		msgList  chan Message
		userInfo UserInfo
	}

	Handler interface {
		Execute(event slack.Event, botUser UserInfo) ([]slack.Message, error)
		Start() error
		Name() string
	}

	UserInfo interface {
		ID() string
		Name() string
	}

	Event interface {
		Name() string
		Data() interface{}
	}

	SlackClient interface {
		Listen() chan slack.Event
		Send(slack.Message)
		UserInfo() (*slack.UserInfo, error)
	}

	Message interface {
		Text() string
		Channel() string
		User() slack.User
	}
)

const messageBufferSize = 1024

func New(slack SlackClient, handlers []Handler) *Bot {
	return &Bot{
		handlers: handlers,
		slack:    slack,
		msgList:  make(chan Message, messageBufferSize),
	}
}

func (b *Bot) Run(ctx context.Context) {
	go b.handleMessages(ctx)

	ticker := time.NewTicker(time.Minute * 5)
	go b.updateUserInfo(ctx, ticker)

	for event := range b.slack.Listen() {
		for _, handler := range b.handlers {
			go b.execHandler(ctx, event, handler)
		}
	}
}

func (b *Bot) updateUserInfo(ctx context.Context, ticker *time.Ticker) {
	for ; true ; <-ticker.C {
		userInfo, err := b.slack.UserInfo()
		if err != nil {
			usbCtx.ErrLogger(ctx).Printf("UpdateUserInfo error: %s", err)
		}
		b.userInfo = userInfo
	}
}

func (b *Bot) handleMessages(ctx context.Context) {
	for msg := range b.msgList {
		b.slack.Send(msg)
	}
}

func (b *Bot) execHandler(ctx context.Context, event slack.Event, handler Handler) {
	botUserInfo, err := b.UserInfo()
	if err != nil {
		usbCtx.ErrLogger(ctx).Printf("%s handler error: %s event: %+v", handler.Name(), err, event)
	}

	msgList, err := handler.Execute(event, botUserInfo)
	if err != nil {
		usbCtx.ErrLogger(ctx).Printf("%s handler error: %s event: %+v", handler.Name(), err, event)
	}

	b.sendMessage(msgList)
}

func (b *Bot) sendMessage(msgList []slack.Message) {
	for _, msg := range msgList {
		b.msgList <- msg
	}
}

func (b *Bot) UserInfo() (UserInfo, error) {
	if b.userInfo == nil {
		return nil, errors.New("no user info collected yet")
	}
	return b.userInfo, nil
}

