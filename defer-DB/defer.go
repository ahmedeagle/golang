package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmedeagle/golang/defer-DB/db"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	database := db.Connect()

	//resource mangment Ensure the database connection is closed during shutdown
	defer func() {
		sqlDB, err := database.DB()
		if err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Fatalf("Failed to close database: %v", err)
			}
			log.Println("Database connection closed")
		}
	}()

	// Graceful shutdown handling
	go func() {
		if err := r.Run("127.0.0.1:8080"); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	log.Println("Server is running...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server gracefully...")
}
