package main

import (
	"errors"
	"goMuAPI/main/db"
)

type CreateSong struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

func NewSong(title, artist, genre string, year int) (*db.Song, error) {

	if len(title) == 0 || len(title) > 100 {
    return nil, errors.New("title must be between 1 and 100 characters")
  }

	if len(artist) == 0 || len(artist) > 50 {
    return nil, errors.New("artist name must be between 1 and 50 characters")
  }

    // Validate the genre
  if len(genre) == 0 || len(genre) > 50 {
    return nil, errors.New("genre must be between 1 and 50 characters")
  }

  if year < 1850 || year > 2030 {
    return nil, errors.New("invalid year. It must be a positive integer between 1850 and 2030")
  }

	return &db.Song{
		Title:  title,
		Artist: artist,
		Genre:  genre,
		Year:   year,
	}, nil
}
