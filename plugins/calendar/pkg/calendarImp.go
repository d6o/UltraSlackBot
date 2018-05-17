package pkg

import (
	"time"
	"os"
	"io/ioutil"
	"encoding/csv"
	"bytes"
	"io"
	"strings"
	"strconv"
	"math/rand"
)

const DB_FILE = "calendar_db.csv"
const DB_COLUMNS = "id,date,title,description\n"
const DB_COLUMNS_TITLE = "#\tDate\tTitle\tDesc."

type Calendar interface {
	Add(date time.Time, title, description string) (string, error)
	List(date time.Time) (string, error)
	Remove(id int) (string, error)
}

type CalendarImp struct {

}

func (calendar *CalendarImp) Add(date time.Time, title, description string) (string, error) {
	retStr := ""
	retStr += "date = " + date.Format("02/01/2006")
	retStr += "title = " + title
	retStr += "description = " + description

	createDbIfNeeded()
	file, err := os.OpenFile(DB_FILE, os.O_APPEND, 0755)
	if err != nil {
		return "", err
	}
	writer := csv.NewWriter(file)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int()
	record := []string{
		strconv.Itoa(id),
		date.Format("02/01/2006"),
		title,
		description,
	}

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
		record, err := reader.Read()
		if err == io.EOF {
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

func (calendar *CalendarImp) Remove(id int) (string, error) {
	reader, err := calendar.readDb()
	if err != nil {
		return "", err
	}
	file, err := os.OpenFile(DB_FILE, os.O_RDWR, 0755)
	if err != nil {
		return "", err
	}
	writer := csv.NewWriter(file)
	header, err := reader.Read()
	if err != nil {
		return "", err
	}
	writer.Write(header)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break;
		}
		if err != nil {
			return "", err
		}
		recId, err := strconv.Atoi(record[0])
		if err != nil {
			return "", err
		}
		if recId != id {
			error := writer.Write(record)
			if error != nil {
				return "", error
			}
		}
	}
	return "Entry removed", nil
}

func createDbIfNeeded()  error {
	_, error := os.Open(DB_FILE)
	if error == nil || !os.IsNotExist(error) {
		return nil, error
	}
	file, error := os.Create(DB_FILE)
	if error != nil {
		return nil, error
	}
	//_, error = file.WriteString(DB_COLUMNS)
	//return file, error
}

func openDbForWrite() (*os.File, error) {
	err := createDbIfNeeded()
	if err != nil {
		return nil, err
	}
	return os.OpenFile(DB_FILE, os.O_RDWR, 0755)
}

func (calendar *CalendarImp) readDb() (*csv.Reader, error) {
	createDbIfNeeded()
	data, error := ioutil.ReadFile(DB_FILE)
	if error != nil {
		return nil, error
	}
	return csv.NewReader(bytes.NewReader(data)), nil
}