package recognition

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/dbutil"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
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

				// check existance

				err := db.Transaction(func(tx *gorm.DB) error {

					var idn IdentifiedSong
					result := tx.Model(&IdentifiedSong{}).Where(&IdentifiedSong{IdInRecognizeService: match.MatchedSongId}).Find(&idn).Limit(1)
					if result.Error != nil {
						return fmt.Errorf("unexpected error on exist check : %v", err)
					}

					if result.RowsAffected == 0 {

						idn = IdentifiedSong{
							Meta:                 match.MatchedSongMeta,
							CreatedAt:            time.Now(),
							IdInRecognizeService: match.MatchedSongId,
							Name:                 match.SongName,
						}

						err := tx.Create(&idn).Error
						if err != nil {
							return fmt.Errorf("expected error on insert idn: %v", err)
						}
					}

					err = db.Raw("UPDATE users_songs SET identified_song_id = ?, is_identified = 1 WHERE id = ?", idn.ID, song.ID).Error
					if err != nil {
						return fmt.Errorf("unexpected error on update users_songs")
					}

					return nil
				})

				if err != nil {
					log.Fatalf("unexpected error on transaction: %v", err)
					return
				}

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
