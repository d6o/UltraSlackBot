package cmd

import (
	"../pkg"
	"github.com/akamensky/argparse"
	"fmt"
	"time"
	"errors"
	"strconv"
)


type ArgParser struct {
	Calendar pkg.Calendar
}

func (argParser ArgParser) ParseCmd(args []string) (string, error) {
	retString := ""
	parser := argparse.NewParser("Calendar", "Keep an event calendar for your chat!")
	list := parser.NewCommand("list", "List events")
	//list.Usage("list [DATE]\nDATE in format DD/MM/YYYY")
	list_date := list.String("d", "date", &argparse.Options{Required: false, Help: "Max date"})

	add := parser.NewCommand("add", "Add new event")
	//add.Usage("add DATE TITLE \nDATE in format DD/MM/YYYY")
	add_date := add.String("d", "date", &argparse.Options{Required: true, Help: "Event date"})
	add_title := add.String("t", "title", &argparse.Options{Required: true, Help: "Event title"})
	add_description := add.String("de", "description", &argparse.Options{Required: false, Help: "Event description"})

	remove := parser.NewCommand("remove", "Remove event")
	//remove.Usage("remove ID")
	remove_id := remove.String("i", "id", &argparse.Options{Required: true, Help: "Event entry Id"})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	if list.Happened() {
		t, err := time.Parse("02/01/2006", *list_date)
		if err != nil {
			return "", err
		}
		retString, err = argParser.Calendar.List(t)
	}
	if add.Happened() {
		t, err := time.Parse("02/01/2006", *add_date)
		if err != nil {
			return "", err
		}
		if len(*add_title) == 0 {
			return "",errors.New("Invalid title")
		}
		retString, err = argParser.Calendar.Add(t, *add_title, *add_description)
	}
	if remove.Happened() {
		i, err := strconv.Atoi(*remove_id)
		if err != nil {
			return "",err
		}
		retString, err = argParser.Calendar.Remove(i)
	}
	return retString, nil
}