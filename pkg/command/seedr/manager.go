package seedr

import (
	"fmt"
	"github.com/disiqueira/ultraslackbot/pkg/seedr"
	"strings"
)

const (
	streamURL = "http://github.freevpnforlife.tk/player/#%s"
)

type (
	seedrManager struct {
		client *seedr.Seedr
	}
)

func newSeedrManager(client *seedr.Seedr) *seedrManager {
	return &seedrManager{
		client: client,
	}
}

func (s *seedrManager) Folders() (string, error) {
	folders, err := s.client.Folders()
	fmt.Printf("%+v\n", folders)
	fmt.Printf("%+v\n", err)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", folders.Folders)

	var msgList []string
	for _, folder := range folders.Folders {
		msgList = append(msgList, fmt.Sprintf("%d - %s", folder.ID, folder.Name))
	}

	return strings.Join(msgList, "\n"), nil
}

func (s *seedrManager) Folder(id int) (string, error) {
	folder, err := s.client.Folder(id)
	if err != nil {
		return "", err
	}

	var msgList []string
	for _, file := range folder.Files {
		msgList = append(msgList, fmt.Sprintf("%d - %s", file.ID, file.Name))
	}

	return strings.Join(msgList, "\n"), nil
}

func (s *seedrManager) HLS(id int) (string, error) {
	url, err := s.client.HLS(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(streamURL, url), nil
}
