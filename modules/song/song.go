package song

import (
	"context"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/dbutil"
	"gorm.io/gorm"
)

func insertSongLog(ctx context.Context, song UsersSong, durationPlayed int) error {

	db, err := dbutil.GormDB(ctx)
	if err != nil {
		return err
	}

	var dbsong UsersSong
	// check existance of song in users_songs
	err = db.Model(&UsersSong{}).Where(&UsersSong{
		UserId: song.UserId,
		SongId: song.SongId,
	}).Take(&dbsong).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	var recordId int64 = dbsong.ID
	if err == gorm.ErrRecordNotFound {
		// insert to users_songs table

		song.CreatedAt = time.Now()
		err = db.Create(&song).Error
		if err != nil {
			return err
		}

		recordId = song.ID

	}
	// add song log
	log := SongLog{
		SongId:         recordId,
		UserId:         song.UserId,
		DurationPlayed: durationPlayed,
		CreatedAt:      time.Now(),
	}

	err = db.Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}
