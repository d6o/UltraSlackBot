package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"

	cnpjHelper "github.com/martinusso/go-docs/cnpj"
)

const (
	pattern = "(?i)^(cnpj)"
	space = " "
	msgPrefix = "Here's your CNPJ: %s"
	generateCommand = "generate"
	validateCommand = "validate"
	validCNPJ                = "%s is a valid CNPJ"
	invalidCNPJ              = "%s is NOT a valid CNPJ"
	invalidParams = "Invalid parameters."
)

type (
	cnpj struct {
		plugin.BasicCommand
	}
)

func (c *cnpj) Name() string {
	return "cnpj"
}

func (c *cnpj) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *cnpj) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *cnpj) command(text string) (string, error) {
	args := strings.Split(text, space)
	if len(args) == 1 || (len(args) >= 1 && args[1] == generateCommand) {
		return c.generate()
	}

	if len(args) >= 3 && args[1] == validateCommand {
		return c.validate(args[2])
	}

	return invalidParams, nil
}

func (c *cnpj) generate() (string, error) {
	return fmt.Sprintf(msgPrefix, cnpjHelper.Generate()), nil
}

func (c *cnpj) validate(cnpj string) (string, error) {
	if cnpjHelper.Valid(cnpj) {
		return fmt.Sprintf(validCNPJ, cnpj), nil
	}

	return fmt.Sprintf(invalidCNPJ, cnpj), nil
}

var CustomPlugin cnpj
