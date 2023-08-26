package song

import (
	"database/sql"
	"time"
)

type Song struct {
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
}

type UsersSong struct {
	ID           int64
	SongId       int64
	UserId       int64
	Title        string
	Year         int
	Duration     int64
	Data         string
	DateModified int64
	AlbumId      int64
	AlbumName    string
	ArtistId     int64
	ArtistName   string
	Composer     sql.NullString
	AlbumArtist  sql.NullString
	CreatedAt    time.Time
}