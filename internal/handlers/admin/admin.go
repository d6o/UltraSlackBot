package admin

import (
	"fmt"
	"strings"

	"context"

	"github.com/disiqueira/ultraslackbot/internal/bot"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

type (
	Admin struct {
		bot *bot.Bot
	}
)

func New(bot *bot.Bot) *Admin {
	return &Admin{
		bot: bot,
	}
}

func (a *Admin) Execute(ctx context.Context, message slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(message)
	if err != nil {
		return nil, nil
	}
	return a.handleMessageEvent(msg, botUser)
}

func (a *Admin) Name() string {
	return "admin"
}

func (a *Admin) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")
	if len(args) < 2 || args[0] != a.Name() {
		return nil, nil
	}

	var messages []string
	switch args[1] {
	case "list":
		messages = a.list()
	case "disable":
		messages = a.disable(args[2:])
	case "enable":
		messages = a.enable(args[2:])
	}

	var outMessages []slack.Message
	for _, m := range messages {
		outMessages = append(outMessages, slack.NewMessage(m, messageEvent.Channel(), botUser))

	}
	return outMessages, nil
}

func (a *Admin) list() []string {
	messages := make([]string, 0, len(a.bot.Handlers()))

	for _, handler := range a.bot.Handlers() {
		m := fmt.Sprintf("%s: %t", handler.Name(), handler.Enabled())
		messages = append(messages, m)
	}
	return messages
}

func (a *Admin) enable(args []string) []string {
	if len(args) < 1 {
		return nil
	}
	handler := args[0]
	a.ensureHandlerStatus(true, handler)
	return []string{
		fmt.Sprintf("%s is now enabled", handler),
	}
}

func (a *Admin) disable(args []string) []string {
	if len(args) < 1 {
		return nil
	}

	handler := args[0]
	if handler == "admin" {
		return []string{
			fmt.Sprintf("%s can not be disabled", handler),
		}
	}

	a.ensureHandlerStatus(false, handler)
	return []string{
		fmt.Sprintf("%s is now disabled", handler),
	}
}

func (a *Admin) ensureHandlerStatus(enabled bool, handler string) error {
	h, ok := a.bot.Handlers()[strings.ToLower(handler)]
	if !ok {
		return fmt.Errorf("handler %s is not loaded", handler)
	}

	if h.Enabled() == enabled {
		return nil
	}
	if !enabled {
		h.Disable()
		return nil
	}

	h.Enable()
	return nil
}
