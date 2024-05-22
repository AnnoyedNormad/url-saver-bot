package main

import (
	"fmt"
	"log"
	"url-saver-bot/internal/config"
)

func main() {
	var token string

	fmt.Println(token)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: create logger

	//TODO: create fetcher

	//TODO: create processor

	//TODO: create consumer
}
