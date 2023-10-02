package main

import (
	"fmt"
	"goMuAPI/main/db"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn := db.ConnectDB(os.Getenv("DSN"))

	s := db.SongService{DB: conn}
	s.CreateSong(
		db.SongData{
			Title: "Metalnigus",
			Artist: "Edge",
			Genre: "Metal",
			Year: 2012,
		},
	)
	fmt.Printf("Api is now working!!!")
}