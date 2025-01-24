package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	gormConfig := gorm.Config{}

	// create dsn

	dsn := fmt.Sprintf("username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4,utf8&parseTime=True")
	if dsn == "" {
		return nil
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	if gin.IsDebugging() {
		gormConfig = gorm.Config{
			Logger: newLogger,
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
