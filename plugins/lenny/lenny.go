package main

import (
	"math/rand"
	"regexp"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern = "(?i)\\b(lenny)\\b"
)

var (
	lennyList = []string{
		"( \u0361\u00b0 \u035c\u0296 \u0361\u00b0)",
		"( \u0360\u00b0 \u035f\u0296 \u0361\u00b0)",
		"\u1566( \u0361\u00b0 \u035c\u0296 \u0361\u00b0)\u1564",
		"( \u0361\u00b0 \u035c\u0296 \u0361\u00b0)",
		"( \u0361~ \u035c\u0296 \u0361\u00b0)",
		"( \u0361o \u035c\u0296 \u0361o)",
		"\u0361\u00b0 \u035c\u0296 \u0361 -",
		"( \u0361\u0361 \u00b0 \u035c \u0296 \u0361 \u00b0)\ufeff",
		"( \u0361 \u0361\u00b0 \u0361\u00b0  \u0296 \u0361\u00b0 \u0361\u00b0)",
		"(\u0e07 \u0360\u00b0 \u035f\u0644\u035c \u0361\u00b0)\u0e07",
		"( \u0361\u00b0 \u035c\u0296 \u0361 \u00b0)",
		"( \u0361\u00b0\u256d\u035c\u0296\u256e\u0361\u00b0 )",
		"(    \u0361\u00b0 \u035c  \u0361\u00b0    )",
		"( \u0361\u00b0                  \u035c                      \u0361\u00b0 )",
		"(\u0e07     \u0360\u00b0 \u035f   \u0361\u00b0    )\u0e07",
		"(    \u0361\u00b0_ \u0361\u00b0    )",
		"(\ufffd    \u0361\u00b0 \u035c  \u0361\u00b0    )\ufffd",
		"(   \u25d5  \u035c  \u25d5   )",
		"(   \u0361~  \u035c   \u0361\u00b0   )",
		"(    \u0360\u00b0 \u035f   \u0361\u00b0    )",
		"(   \u0ca0  \u035c  \u0ca0   )",
		"(    \u0ca5  \u035c  \u0ca5    )",
		"(    \u0361^ \u035c  \u0361^    )",
		"(    \u0ca5 _  \u0ca5    )",
		"(    \u0361\u00b0 \uff0d \u0361\u00b0    )",
		"\u2570(      \u0361\u00b0  \u035c   \u0361\u00b0)\u2283\u2501\u2606\u309c\u30fb\u3002\u3002\u30fb\u309c\u309c\u30fb\u3002\u3002\u30fb\u309c\u2606\u309c\u30fb\u3002\u3002\u30fb\u309c\u309c\u30fb\u3002\u3002\u30fb\u309c",
		"\u2534\u252c\u2534\u252c\u2534\u2524(    \u0361\u00b0 \u035c  \u251c\u252c\u2534\u252c\u2534\u252c",
		"(    \u2310\u25a0 \u035c   \u25a0  )",
		"(    \u0361~ _ \u0361~    )",
		"@=(   \u0361\u00b0 \u035c  \u0361\u00b0  @ )\u2261",
		"(    \u0361\u00b0\u06a1 \u0361\u00b0    )",
		"(  \u2716_\u2716  )",
		"(\u3065    \u0361\u00b0 \u035c  \u0361\u00b0    )\u3065",
		"\u10da(   \u0361\u00b0 \u035c  \u0361\u00b0   \u10da)",
		"(    \u25c9 \u035c  \u0361\u25d4    )",
	}
)

type (
	lenny             struct{
		plugin.BasicCommand
	}
)

func (d *lenny) Start(specs bot.Specs) error {
	return nil
}

func (d *lenny) Name() string {
	return "lenny"
}

func (d *lenny) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return d.HandleEvent(event, botUser, d.matcher, d.command)
}

func (d *lenny) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (d *lenny) command(text string) (string, error) {
	return lennyList[rand.Intn(len(lennyList))], nil
}

var CustomPlugin lenny
