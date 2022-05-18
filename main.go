package main

import (
	"github.com/G-Basilio/goApiRest/database"
	"github.com/G-Basilio/goApiRest/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
