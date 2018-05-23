package cmd

import (
	"github.com/disiqueira/ultraslackbot/plugins/calendar/pkg"
	"time"
	"errors"
	"log"
)

const (
	CMD_LIST = "list"
	CMD_ADD = "add"
	CMD_REMOVE = "remove"
	MSG_NO_CMD = "Use one of the commands: '" + CMD_LIST + "', '" + CMD_ADD + "' or '" + CMD_REMOVE
	MSG_LIST = "Usage: list [DD/MM/YYYY]"
	MSG_ADD = "Usage: add DD/MM/YYYY Title [Description]"
	MSG_REMOVE = "Usage: remove TITLE"
)

type ArgParser struct {
	Calendar pkg.Calendar
}

func (argParser ArgParser) ParseCmd(args []string) (string, error) {
	log.Print("ParseCmd()")
	if len(args) == 0 || args[0] == "" {
		return "", errors.New(MSG_NO_CMD)
	}
	command := args[0]
	parameters := args[1:len(args)]
	switch command {
	case CMD_LIST:
		return argParser.list(parameters)
	case CMD_ADD:
		return argParser.add(parameters)
	case CMD_REMOVE:
		return argParser.remove(parameters)
	default:
		return "", errors.New("Invalid command. " + MSG_NO_CMD)
	}
}

func (argParser ArgParser) list(args []string) (string, error) {
	log.Print("list()")
	var date time.Time
	if (len(args) > 0) {
		if (args[0] == "-h" || args[0] == "--help") {
			return MSG_LIST, nil
		}
		var err error
		date, err = time.Parse("02/01/2006", args[0])
		if err != nil {
			return "", errors.New("Invalid parameter date: " + err.Error() +
				"\n" + MSG_LIST)
		}
	}
	return argParser.Calendar.List(date)
}

func (argParser ArgParser) add(args []string) (string, error) {
	log.Print("add()")
	if len(args) < 2 {
		return "", errors.New("Invalid parameters. \n" + MSG_ADD)
	}
	t, err := time.Parse("02/01/2006", args[0])
	if err != nil {
		return "", errors.New(err.Error() + "\n" + MSG_ADD)
	}
	if len(args[1]) == 0 || args[1] == "" {
		return "",errors.New("Invalid title. \n" + MSG_ADD)
	}
	desc := ""
	if len(args) >= 3 {
		desc = args[2]
	}
	return argParser.Calendar.Add(t, args[1], desc)
}

func (argParser ArgParser) remove(args []string) (string, error) {
	log.Print("remove()")
	if len (args) < 1 {
		return "", errors.New("Invalid parameters. \n" + MSG_REMOVE)
	}
	return argParser.Calendar.Remove(args[0])
}