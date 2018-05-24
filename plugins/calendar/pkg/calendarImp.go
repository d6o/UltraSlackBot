package pkg

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	dbFile         = "calendar_db.csv"
	dbColumns      = "date,title,description\n"
	dbColumnsTitle = "Date\tTitle\tDesc."
	dbFileMode     = 0755
)

type Calendar interface {
	Add(date time.Time, title, description string) (string, error)
	List() (string, error)
	Remove(title string) (string, error)
}

type CalendarImp struct {
}

func (calendar *CalendarImp) Add(date time.Time, title, description string) (string, error) {
	createDbIfNeeded()
	file, err := os.OpenFile(dbFile, os.O_WRONLY|os.O_APPEND, dbFileMode)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		date.Format("02/01/2006"),
		title,
		description,
	}
	log.Print("CalendarImp.Add(): record = ")
	log.Print(record)

	err = writer.Write(record)
	if err != nil {
		return "", err
	}
	return "Event added!", nil
}

func (calendar *CalendarImp) List() (string, error) {
	retStr := dbColumnsTitle
	reader, err := calendar.readDb()
	if err != nil {
		return "", err
	}
	for {
		log.Print("List(): reading record")
		record, err := reader.Read()
		if err == io.EOF {
			log.Print("List(): EOF")
			break
		}
		if err != nil {
			return "", err
		}
		retStr += "\n"
		retStr += strings.Join(record, "\t")
	}
	return retStr, nil
}

func (calendar *CalendarImp) Remove(title string) (string, error) {
	log.Print("Remove()")
	reader, err := calendar.readDb()
	if err != nil {
		return "", err
	}
	file, err := os.OpenFile(dbFile, os.O_RDWR, 0755)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	found := false
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if record[1] != title {
			error := writer.Write(record)
			if error != nil {
				return "", error
			}
			log.Print("'" + record[1] + "' != '" + title + "'")
		} else {
			found = true
		}
	}
	if found {
		return "Entry removed", nil
	}
	return "", errors.New("Entry not found")
}

func createDbIfNeeded() error {
	log.Print("createDbIfNeeded()")
	file, err := os.Open(dbFile)
	if err == nil {
		file.Close()
		return err
	}
	log.Print("createDbIfNeeded(): will create db")
	file, err = os.Create(dbFile)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (calendar *CalendarImp) readDb() (*csv.Reader, error) {
	log.Print("readDb()")
	createDbIfNeeded()
	data, error := ioutil.ReadFile(dbFile)
	if error != nil {
		return nil, error
	}
	return csv.NewReader(bytes.NewReader(data)), nil
}
