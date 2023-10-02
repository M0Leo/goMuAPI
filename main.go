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
	fmt.Printf("Api is now working!!!")
}