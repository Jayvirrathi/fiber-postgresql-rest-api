package database

import (
	"fmt"
	// "os"
	"strconv"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



//   os.Setenv("DB_PORT", "5432")
//   os.Setenv("DB_USER", "postgres")
//   os.Setenv("DB_PASSWORD", "1234")
//   os.Setenv("DB_HOST", "localhost")
//   os.Setenv("DB_NAME", "test")


var (
	
	// DBConn is a pointer to gorm.DB
	DBConn   *gorm.DB
	user     = "postgres"
	password = "1234"
	host     = "localhost"
	db       = "postgres"
	port     = "5433"
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// host     = os.Getenv("DB_HOST")
	// db       = os.Getenv("DB_NAME")
	// port     = os.Getenv("DB_PORT")
)




// Connect creates a connection to database
func Connect() (err error) {
	port, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, port)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := DBConn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
