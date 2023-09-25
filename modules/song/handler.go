package song

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AhmadRezaZarei/musicrecommander/modules/ginutil"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {

	r.POST("/song/log", func(ctx *gin.Context) {

		userId := 1
		file, err := ctx.FormFile("file")
		if err == nil {
			// TODO unique name
			ctx.SaveUploadedFile(file, filepath.Join("songs", file.Filename+".mp3"))
		}

		req := ctx.Request
		songId, err := strconv.Atoi(req.FormValue("id"))

		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		title := req.FormValue("title")

		year, err := strconv.Atoi(req.FormValue("year"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		duration, err := strconv.Atoi(req.FormValue("duration"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		date := req.FormValue("date")

		albumId, err := strconv.Atoi(req.FormValue("album_id"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		albumName := req.FormValue("album_name")
		composer := req.FormValue("composer")

		artistId, err := strconv.Atoi(req.FormValue("artist_id"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}
		artistName := req.FormValue("artist_name")
		albumArtist := req.FormValue("album_artist")

		songStartedAt, err := strconv.Atoi(req.FormValue("song_started_at"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		songEndedAt, err := strconv.Atoi(req.FormValue("song_ended_at"))
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		userSong := UsersSong{
			SongId:     int64(songId),
			UserId:     int64(userId),
			Title:      title,
			Year:       year,
			Duration:   int64(duration),
			Data:       date,
			AlbumId:    int64(albumId),
			AlbumName:  albumName,
			ArtistId:   int64(artistId),
			ArtistName: artistName,
			Composer: sql.NullString{
				String: composer,
				Valid:  true,
			},
			AlbumArtist: sql.NullString{
				String: albumArtist,
				Valid:  true,
			},
			CreatedAt: time.Now(),
		}

		err = insertSongLog(ctx, &userSong, songEndedAt-songStartedAt)
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": nil,
		})
	})

	r.POST("/song/logs", func(ctx *gin.Context) {
		userId := 1

		var req SongLogsRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		err = insertSongLogs(ctx, int64(userId), req.Logs)
		if err != nil {
			ginutil.SendWrappedInternalServerError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"error": nil,
		})

	})

}
