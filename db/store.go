package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SongService interface {
	CreateSong(*Song) error
	GetSongByID(id int) (*Song, error)
	GetSongs() ([]*Song, error)
}

type MySQLStore struct {
	db *gorm.DB
}

type Song struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

func NewMySQLStore(dsn string) (*MySQLStore, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Song{})
	return &MySQLStore{
		db: db,
	}, err
}

func (store *MySQLStore) CreateSong(song *Song) error {
	result := store.db.Create(song)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (store *MySQLStore) GetSongByID(id int) (*Song, error) {
	var song Song
	if err := store.db.First(&song, id).Error; err != nil {
		return nil, err
	}

	return &song, nil
}

func (store *MySQLStore) GetSongs() ([]*Song, error) {
	var songs []*Song
	if err := store.db.Find(&songs).Error; err != nil {
		return nil, err
	}

	return songs, nil
}
