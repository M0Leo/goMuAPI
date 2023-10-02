package db

import "gorm.io/gorm"

type SongService struct {
	DB *gorm.DB
}

type Song struct {

	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`

}

type SongData struct {
	Title string
	Artist string
	Genre string
	Year int
}

func (s *SongService) CreateSong(data SongData) error {
	song := Song{
		Title:  data.Title,
		Artist: data.Artist,
		Genre:  data.Genre,
		Year:   data.Year,
	}

	result := s.DB.Create(&song)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

