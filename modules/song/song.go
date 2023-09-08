package song

import (
	"context"
	"database/sql"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/dbutil"
	"gorm.io/gorm"
)

func insertSongLog(ctx context.Context, song *UsersSong, durationPlayed int) error {

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

func insertSongLogs(ctx context.Context, userId int64, logs []*RequestSongLog) error {

	for _, log := range logs {

		composer := sql.NullString{
			Valid: false,
		}
		if log.Composer != nil {
			composer = sql.NullString{
				String: *log.Composer,
				Valid:  true,
			}
		}

		albumArtist := sql.NullString{
			Valid: false,
		}
		if log.AlbumArtist != nil {
			albumArtist = sql.NullString{
				String: *log.AlbumArtist,
				Valid:  true,
			}
		}

		err := insertSongLog(ctx, &UsersSong{
			SongId:      log.ID,
			UserId:      userId,
			Title:       log.Title,
			Year:        log.Year,
			Duration:    log.Duration,
			Data:        log.Data,
			AlbumId:     log.AlbumId,
			AlbumName:   log.AlbumName,
			ArtistId:    log.ArtistId,
			ArtistName:  log.ArtistName,
			Composer:    composer,
			AlbumArtist: albumArtist,
			CreatedAt:   time.Now(),
		}, int(log.SongEndedAt)-int(log.SongStartedAt))
		if err != nil {
			return err
		}

	}

	return nil
}
