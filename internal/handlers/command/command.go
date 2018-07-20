package command

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"github.com/disiqueira/ultraslackbot/internal/bot"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	usbctx "github.com/disiqueira/ultraslackbot/internal/context"
)

type (
	Command struct {
		cmd    *cobra.Command
		sender *sender
	}

	sender struct {
		msgList []slack.Message
		channel string
		user    slack.User
	}
)

const (
	suffix = "!"
)

func New(cmdList []*cobra.Command) *Command {
	c := &Command{
		cmd: &cobra.Command{
			Use:   "!",
			Short: "The best slack bot in the world.",
			SilenceErrors: false,
			SilenceUsage: false,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		},
		sender: &sender{},
	}

	c.cmd.AddCommand(cmdList...)
	return c
}

func (c *Command) Name() string {
	return "Command"
}

func (c *Command) Execute(ctx context.Context, event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return c.handleMessageEvent(ctx, msg, botUser)
}

func (c *Command) handleMessageEvent(ctx context.Context, messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	usbctx.OutLogger(ctx).Printf("%+v", messageEvent)

	msg := strings.TrimSpace(messageEvent.Text())
	if len(msg) == 0 || string(msg[0]) != suffix {
		return nil, nil
	}

	c.cmd.SetArgs(strings.Fields(string(msg[1:])))
	c.cmd.SetOutput(c.sender)

	c.sender.channel = messageEvent.Channel()
	c.sender.user = botUser
	c.sender.msgList = make([]slack.Message, 0)

	err := c.cmd.Execute()
	if err != nil {
		c.sender.Write([]byte(err.Error()))
	}

	return c.sender.msgList, err
}

func (s *sender) Write(p []byte) (n int, err error) {
	s.msgList = append(s.msgList, slack.NewMessage(string(p), s.channel, s.user))
	return len(p), nil
}
