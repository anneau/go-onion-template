package main

import (
	"github.com/anneau/go-template/api/server"
	"github.com/anneau/go-template/config"
	"github.com/anneau/go-template/infra/database"
)

func main() {
	config := config.NewConfig()

	dbConn, err := database.NewConnection(&config.Database)

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	s, err := server.NewServer(&config.HTTPServer, dbConn)

	if err != nil {
		panic(err)
	}

	s.Run()
}
