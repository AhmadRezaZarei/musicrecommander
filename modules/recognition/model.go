package recognition

import "time"

type IdentifiedSong struct {
	ID        int64
	name      string
	meta      string
	CreatedAt time.Time
}

type RecognitionResult struct {
	AlignTime       string       `json:"align_time"`
	FingerprintTime string       `json:"fingerprint_time"`
	QueryTime       string       `json:"query_time"`
	Results         []*MatchSong `json:"results"`
	TotalTime       string       `json:"total_time"`
}

type MatchSong struct {
	FileSha1                string `json:"file_sha1"`
	FingerprintedConfidence string `json:"fingerprinted_confidence"`
	FingerprintedHashesInDb string `json:"fingerprinted_hashes_in_db"`
	HashesMatchedInInput    string `json:"hashes_matched_in_input"`
	InputConfidence         string `json:"input_confidence"`
	InputTotalHashes        string `json:"input_total_hashes"`
	Offset                  string `json:"offset"`
	OffsetSeconds           string `json:"offset_seconds"`
	SongID                  string `json:"song_id"`
	SongMeta                string `json:"song_meta"`
	SongName                string `json:"song_name"`
}

type BaseError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
type RecognizeResponse struct {
	Result RecognitionResult `json:"result"`
	Error  *BaseError        `json:"error"`
}

type SongMatchResult struct {
	IsMatched       bool
	MatchedSongMeta string
	MatchedSongId   string
}
