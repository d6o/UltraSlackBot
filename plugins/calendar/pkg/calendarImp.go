package pkg

import (
	"time"
	"os"
	"io/ioutil"
	"encoding/csv"
	"bytes"
	"io"
	"strings"
	"errors"
	"log"
)

const DB_FILE = "calendar_db.csv"
const DB_COLUMNS = "date,title,description\n"
const DB_COLUMNS_TITLE = "Date\tTitle\tDesc."

type Calendar interface {
	Add(date time.Time, title, description string) (string, error)
	List(date time.Time) (string, error)
	Remove(title string) (string, error)
}

type CalendarImp struct {

}

func (calendar *CalendarImp) Add(date time.Time, title, description string) (string, error) {
	createDbIfNeeded()
	file, err := os.OpenFile(DB_FILE, os.O_WRONLY|os.O_APPEND, 0755)
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

func (calendar *CalendarImp) List(date time.Time) (string, error){
	retStr := DB_COLUMNS_TITLE
	// retStr += "date = " + date.Format("02/01/2006")
	reader, err := calendar.readDb()
	if err != nil {
		return "", err
	}
	for {
		log.Print("List(): reading record")
		record, err := reader.Read()
		if err == io.EOF {
			log.Print("List(): EOF")
			break;
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
	file, err := os.OpenFile(DB_FILE, os.O_RDWR, 0755)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	/*header, err := reader.Read()
	if err != nil {
		return "", err
	}
	writer.Write(header)*/
	found := false
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break;
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
	} else {
		return "", errors.New("Entry not found")
	}
}

func createDbIfNeeded()  error {
	log.Print("createDbIfNeeded()")
	file, error := os.Open(DB_FILE)
	if error == nil {
		file.Close()
		return error
	}
	log.Print("createDbIfNeeded(): will create db")
	file, error = os.Create(DB_FILE)
	if error != nil {
		return error
	}
	file.Close()
	//_, error = file.WriteString(DB_COLUMNS)
	//return file, error
	return nil
}

func (calendar *CalendarImp) readDb() (*csv.Reader, error) {
	log.Print("readDb()")
	createDbIfNeeded()
	data, error := ioutil.ReadFile(DB_FILE)
	if error != nil {
		return nil, error
	}
	log.Print("readDb(): dump = " + string(data))
	return csv.NewReader(bytes.NewReader(data)), nil
}