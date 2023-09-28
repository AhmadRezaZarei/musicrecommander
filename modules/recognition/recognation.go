package recognition

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/songlog"
	"gorm.io/gorm"
)

func fetchUnidentifiedUserSongs(tx *gorm.DB, limit int) ([]*songlog.UsersSong, error) {

	var userSongs []*songlog.UsersSong
	err := tx.Raw("SELECT * FROM users_songs WHERE is_identified = 0 AND filename IS NOT NULL").Limit(10).Scan(&userSongs).Error

	if err != nil {
		return nil, err
	}

	return userSongs, nil

}

func recognizeSong(filename string) (*SongMatchResult, error) {
	matches, err := callRecognizeService(filename)

	if err != nil {
		return nil, err
	}

	return &SongMatchResult{
		IsMatched:       false,
		MatchedSongMeta: "",
		MatchedSongId:   "",
	}, nil

	// find best match

	maxConf := 0.0
	bestOffset := 0.0
	var matched *MatchSong = nil

	for i, m := range matches {

		conf, err := strconv.ParseFloat(m.InputConfidence, 64)
		if err != nil {
			return nil, err
		}

		offset, err := strconv.ParseFloat(m.Offset, 64)
		if err != nil {
			return nil, err
		}

		if (i == 0) || (conf > maxConf) || (conf == maxConf && offset < bestOffset) {
			maxConf = conf
			bestOffset = offset
			matched = m
		}

	}

	if maxConf > 0.2 {
		return &SongMatchResult{
			IsMatched:       true,
			MatchedSongMeta: matched.SongMeta,
			MatchedSongId:   matched.SongID,
		}, nil
	}

	return &SongMatchResult{
		IsMatched:       false,
		MatchedSongMeta: "",
		MatchedSongId:   "",
	}, nil

}

func callRecognizeService(filename string) ([]*MatchSong, error) {

	fileDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fileName := filename
	filePath := path.Join(fileDir, "songs", fileName)

	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(fw, file); err != nil {
		return nil, err
	}

	w.Close()

	req, err := http.NewRequest("POST", "http://localhost:5678/recognize", &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// decode response

	decoder := json.NewDecoder(resp.Body)

	var response RecognizeResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Result.Results, nil
}
