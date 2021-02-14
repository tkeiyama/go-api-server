package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// DB is db of gorm.
var DB *gorm.DB

// DBConfig represents db configuration.
type DBConfig struct {
	Host     string
	Port     int16
	User     string
	Password string
	DBName   string
}

// BuildDBConfig returns the structure of DBConfig.
func BuildDBConfig() *DBConfig {
	root := os.Getenv("ROOT")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConfig := &DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     root,
		Password: password,
		DBName:   dbName,
	}
	return dbConfig
}

// DbURL shows us DB information.
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
