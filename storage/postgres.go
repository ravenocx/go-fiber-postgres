package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/ritsuhaaa/go-fiber-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMODE  string
}

var config = Config{}

// newconnection -> for connect to database
func ConnectDB(config *Config) (*gorm.DB, error) {

	config.Read()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMODE,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	err = conn.AutoMigrate(
		&models.Books{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return conn, nil

}

func (c *Config) Read() {
	// Read env file
	config.Host = os.Getenv("DB_HOST")
	config.User = os.Getenv("DB_USER")
	config.Password = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.Port = os.Getenv("DB_PORT")
	config.SSLMODE = os.Getenv("DB_SSLMODE")
}
