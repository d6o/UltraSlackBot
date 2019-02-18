package cmd

import (
	"errors"
	"log"
	"time"

	"github.com/disiqueira/ultraslackbot/plugins/calendar/pkg"
)

const (
	cmdList   = "list"
	cmdAdd    = "add"
	cmdRemove = "remove"
	msgList   = cmdList
	msgAdd    = cmdAdd + " DD/MM/YYYY TITLE [DESCRIPTION]"
	msgRemove = cmdRemove + " TITLE"
)

type ArgParser struct {
	Calendar pkg.Calendar
}

var errorNoCmd = errors.New("Use one of these commands: \n" + msgList + "\n" + msgAdd + "\n" + msgRemove)

func (a ArgParser) ParseCmd(args []string) (string, error) {
	log.Print("ParseCmd()")
	if len(args) == 0 || args[0] == "" {
		return "", errorNoCmd
	}
	command := args[0]
	parameters := args[1:]
	switch command {
	case cmdList:
		return a.list(parameters)
	case cmdAdd:
		return a.add(parameters)
	case cmdRemove:
		return a.remove(parameters)
	default:
		return "", errorNoCmd
	}
}

func (a ArgParser) list(args []string) (string, error) {
	log.Print("list()")
	return a.Calendar.List()
}

func (a ArgParser) add(args []string) (string, error) {
	log.Print("add()")
	if len(args) < 2 {
		return "", errors.New("Invalid parameters. \nUsage: " + msgAdd)
	}
	t, err := time.Parse("02/01/2006", args[0])
	if err != nil {
		return "", errors.New(err.Error() + "\nUsage: " + msgAdd)
	}
	if len(args[1]) == 0 || args[1] == "" {
		return "", errors.New("Invalid title. \nUsage: " + msgAdd)
	}
	desc := ""
	if len(args) >= 3 {
		desc = args[2]
	}
	return a.Calendar.Add(t, args[1], desc)
}

func (a ArgParser) remove(args []string) (string, error) {
	log.Print("remove()")
	if len(args) < 1 {
		return "", errors.New("Invalid parameters. \nUsage: " + msgRemove)
	}
	return a.Calendar.Remove(args[0])
}
