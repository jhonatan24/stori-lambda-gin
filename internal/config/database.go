package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"stori-lambda/internal/entity"
	"strconv"
	"sync"
)

type configuration struct {
	Host     string
	Port     int
	Password string
	Dbname   string
	User     string
}

const (
	hostName = "db_hostname"
	port     = "db_port"
	password = "db_password"
	dbName   = "db_name"
	user     = "db_user"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func newConfiguration() *configuration {
	port, _ := strconv.Atoi(os.Getenv(port))
	return &configuration{
		Host:     os.Getenv(hostName),
		Port:     port,
		Password: os.Getenv(password),
		Dbname:   os.Getenv(dbName),
		User:     os.Getenv(user),
	}
}

func DatabaseConnection() *gorm.DB {
	once.Do(func() {
		instance = createConnection()
	})
	return instance
}

func createConnection() *gorm.DB {
	configuration := newConfiguration()
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", configuration.User, configuration.Password, configuration.Host, configuration.Port, configuration.Dbname)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("error: %s", err)
	}
	return db
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&entity.TransactionDetailsEntity{})
	if err != nil {
		fmt.Sprintf("error: %s", err)
	}
}
