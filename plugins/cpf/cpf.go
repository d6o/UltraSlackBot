package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"

	cpfHelper "github.com/martinusso/go-docs/cpf"
)

const (
	pattern = "(?i)^(cpf)"
	space = " "
	msgPrefix = "Here's your CPF: %s"
	generateCommand = "generate"
	validateCommand = "validate"
	validCPF                = "%s is a valid CPF"
	invalidCPF              = "%s is NOT a valid CPF"
	invalidParams = "Invalid parameters."
)

type (
	cpf struct {
		plugin.BasicCommand
	}
)

func (c *cpf) Name() string {
	return "cpf"
}

func (c *cpf) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *cpf) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *cpf) command(text string) (string, error) {
	args := strings.Split(text, space)
	if len(args) == 1 || (len(args) >= 1 && args[1] == generateCommand) {
		return c.generate()
	}

	if len(args) >= 3 && args[1] == validateCommand {
		return c.validate(args[2])
	}

	return invalidParams, nil
}

func (c *cpf) generate() (string, error) {
	return fmt.Sprintf(msgPrefix, cpfHelper.Generate()), nil
}

func (c *cpf) validate(cpf string) (string, error) {
	if cpfHelper.Valid(cpf) {
		return fmt.Sprintf(validCPF, cpf), nil
	}

	return fmt.Sprintf(invalidCPF, cpf), nil
}

var CustomPlugin cpf
