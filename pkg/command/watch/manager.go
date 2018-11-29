package watch

import (
	"fmt"
	"strings"
	"time"

	"github.com/disiqueira/ultraslackbot/pkg/seedr"
	"github.com/disiqueira/ultraslackbot/pkg/yts"
)

const (
	ytsOKStatus = "ok"
	separator   = "#########################"
	streamURL   = "http://github.freevpnforlife.tk/player/#%s"
)

type (
	watchManager struct {
		seedr *seedr.Seedr
		yts   *yts.YTS
	}
)

func newSeedrManager(seedr *seedr.Seedr, yts *yts.YTS) *watchManager {
	return &watchManager{
		seedr: seedr,
		yts:   yts,
	}
}

func (w *watchManager) Watch(q string) (string, error) {
	movies, err := w.yts.Search(q)
	if err != nil {
		return "", fmt.Errorf("error searching for movie: error: %s", err.Error())
	}

	if movies.Status != ytsOKStatus {
		return "", fmt.Errorf("error searching for movie: status: %s %s", movies.Status, movies.StatusMessage)
	}

	if len(movies.Data.Movies) == 0 {
		return "", fmt.Errorf("movie not found. query: %s", q)
	}

	if len(movies.Data.Movies) > 1 {
		var movieList []string

		for _, movie := range movies.Data.Movies {
			movieList = append(movieList, " name: "+movie.TitleLong+"\n"+movie.MediumCoverImage)
		}

		return strings.Join(movieList, separator), nil
	}

	resp, err := w.seedr.Download(movies.Data.Movies[0].Torrents[0].URL)
	if err != nil {
		return "", fmt.Errorf("error downloading movie: error: %s", err.Error())
	}

	if !resp.Result {
		return "", fmt.Errorf("error downloading movie: Code: %d Error: %s", resp.Code, resp.Error)
	}

	time.Sleep(1 * time.Minute)

	folders, err := w.seedr.Folders()
	if err != nil {
		return "", fmt.Errorf("error retriving folders: error: %s", err.Error())
	}

	if !folders.Result {
		return "", fmt.Errorf("error retriving folders: Code: %d", folders.Code)
	}

	words := strings.Split(q, " ")
	var folderID int
	for _, word := range words {
		for _, folder := range folders.Folders {
			if strings.Contains(strings.ToLower(folder.Name), strings.ToLower(word)) {
				folderID = folder.ID
				break
			}
		}
	}

	if folderID == 0 {
		return "", fmt.Errorf("folder not found")
	}

	folder, err := w.seedr.Folder(folderID)
	if err != nil {
		return "", fmt.Errorf("error retriving folder %d: error: %s", folderID, err.Error())
	}

	var finalURL string
	for _, file := range folder.Files {

		if file.StreamVideo {

			finalURL, err = w.seedr.HLS(file.ID)
			if err != nil {
				return "", fmt.Errorf("error retriving HLS for file %d: error: %s", file.ID, err.Error())
			}
			break
		}
	}

	msg := movies.Data.Movies[0].MediumCoverImage + "\n" +
		"Name: " + movies.Data.Movies[0].TitleLong + "\n" +
		"Stream: " + fmt.Sprintf(streamURL, finalURL)

	return msg, nil
}
