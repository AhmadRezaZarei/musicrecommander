package recognition

import (
	"context"
	"fmt"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/dbutil"
	"github.com/go-co-op/gocron"
)

func StartJob() error {

	// steps recognation
	// 1. find max 100 un identified songs
	// send them to recognation service
	// get the result
	// if already exists in out table => make is_identified flag true
	// else make new record on identified_songs table and  make is_identified flag true

	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(20).Seconds().Do(func() {

		db, err := dbutil.GormDB(context.Background())
		if err != nil {
			fmt.Println("unexpected error on create db %v", err)
			return
		}

		unidentifiedSongs, err := fetchUnidentifiedUserSongs(db, 100)

		if len(unidentifiedSongs) == 0 {
			fmt.Println("noting to recognize")
			return
		}

		for _, song := range unidentifiedSongs {

			match, err := recognizeSong(song.Filename.String)
			if err != nil {
				fmt.Println(err.Error())
			}

			if match.IsMatched {
				// insert a record to identified songs table
			}

		}

	})

	if err != nil {
		return err
	}

	s.StartAsync()

	return nil
}
