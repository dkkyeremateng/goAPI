package main

import (
	"log"
	"userAPI/handlers"
	"userAPI/server"

	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

var config = &bongo.Config{
	ConnectionString: "localhost",
	Database:         "umGoAPI",
}

func main() {
	conn, err := bongo.Connect(config)
	if err != nil {
		log.Fatalf("mongodb failed to connet: %v", err)
	}

	r := gin.Default()

	u := handlers.NewHandlers(conn)
	u.SetupRoutes(r)

	srv := server.New(r, ":8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
