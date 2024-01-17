package main

import (
	"exam2/api"
	"exam2/api/handler"
	"exam2/config"
	"exam2/storage/postgres"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}

	defer store.CloseDB()

	handler := handler.New(store)

	api.New(handler)

	fmt.Println("Server is running on port 8088")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("error while running server err:", err.Error())
	}
}
