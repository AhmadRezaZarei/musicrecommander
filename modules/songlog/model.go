package songlog

import (
	"database/sql"
	"time"
)

type Song struct {
	ID           int64   `json:"id"`
	Title        string  `json:"title"`
	Year         int     `json:"year"`
	Duration     int64   `json:"duration"`
	Data         string  `json:"date"`
	DateModified int64   `json:"date_modified"`
	AlbumId      int64   `json:"album_id"`
	AlbumName    string  `json:"album_name"`
	ArtistId     int64   `json:"artist_id"`
	ArtistName   string  `json:"artist_name"`
	Composer     *string `json:"composer"`
	AlbumArtist  *string `json:"album_artist"`
}

type RequestSongLog struct {
	Song
	SongStartedAt int64 `json:"song_started_at"`
	SongEndedAt   int64 `json:"song_ended_at"`
	Timestamp     int64 `json:"timestamp"`
}

type SongLogsRequest struct {
	Logs []*RequestSongLog `json:"logs"`
}

type HistorySong struct {
	ID           int64   `json:"id"`
	Title        string  `json:"title"`
	TrackNumber  int     `json:"track_number"`
	Year         int     `json:"year"`
	Duration     int64   `json:"duration"`
	Data         string  `json:"date"`
	DateModified int64   `json:"date_modified"`
	AlbumId      int64   `json:"album_id"`
	AlbumName    string  `json:"album_name"`
	ArtistId     int64   `json:"artist_id"`
	ArtistName   string  `json:"artist_name"`
	Composer     *string `json:"composer"`
	AlbumArtist  *string `json:"album_artist"`
	TimePlayed   int64   `json:"time_played"`
}

type UsersSong struct {
	ID           int64
	SongId       int64
	UserId       int64
	Title        string
	Year         int
	Duration     int64
	Data         string
	AlbumId      int64
	AlbumName    string
	ArtistId     int64
	ArtistName   string
	Composer     sql.NullString
	AlbumArtist  sql.NullString
	Filename     sql.NullString
	IsIdentified bool
	CreatedAt    time.Time
}

type SongLog struct {
	ID             int64
	UserId         int64
	SongId         int64
	DurationPlayed int
	Timestamp      int
	CreatedAt      time.Time
}
