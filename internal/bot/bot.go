package bot

import (
	"context"
	"errors"
	usbCtx "github.com/disiqueira/ultraslackbot/internal/context"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"time"
	"strings"
)

type (
	Bot struct {
		token    string
		handlers map[string]*HandlerManager
		slack    SlackClient
		msgList  chan Message
		userInfo UserInfo
	}

	HandlerManager struct {
		handler Handler
		enabled bool
	}

	Handler interface {
		Execute(ctx context.Context, event slack.Event, botUser UserInfo) ([]slack.Message, error)
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

	Specs interface {
		All() map[string]interface{}
		Get(k string) (interface{}, bool)
		Set(k, v string)
	}
)

const messageBufferSize = 1024

func New(slack SlackClient) *Bot {
	return &Bot{
		slack:    slack,
		msgList:  make(chan Message, messageBufferSize),
	}
}

func (b *Bot) SetHandlers(handlers []Handler) {
	handlerManagers := map[string]*HandlerManager{}

	for _, h := range handlers {
		handlerManagers[strings.ToLower(h.Name())] = &HandlerManager{
			handler: h,
			enabled: true,
		}
	}

	b.handlers = handlerManagers
}

func (b *Bot) Run(ctx context.Context) {
	go b.handleMessages(ctx)

	t := time.NewTicker(time.Minute * 5)
	go b.updateUserInfo(ctx, t)

	for event := range b.slack.Listen() {
		for _, h := range b.handlers {
			if h.enabled {
				go b.execHandler(ctx, event, h.handler)
			}
		}
	}
}

func (b *Bot) updateUserInfo(ctx context.Context, ticker *time.Ticker) {
	for ; true; <-ticker.C {
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

	msgList, err := handler.Execute(ctx, event, botUserInfo)
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

func (b *Bot) Handlers() map[string]*HandlerManager {
	return b.handlers
}

func (h *HandlerManager) Disable() {
	h.enabled = false
}

func (h *HandlerManager) Enable() {
	h.enabled = true
}

func (h *HandlerManager) Enabled() bool {
	return h.enabled
}

func (h *HandlerManager) Name() string {
	return h.handler.Name()
}
