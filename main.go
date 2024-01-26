package main

import (
	"log"

	"github.com/SergeyMilch/crypto-rate-fetcher/cmd"
	"github.com/joho/godotenv"
)


func init() {
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    cmd.Execute()
}