package main

import (
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

	store, err := db.NewMySQLStore(os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	
	server := NewAPIServer(":8080", store)
	server.Run()
}