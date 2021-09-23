package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Init() error {
	user := viper.GetString("db.user")
	passwd := viper.GetString("db.passwd")
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	database := viper.GetString("db.database")
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, passwd, host, port, database)

	db, err := getDataSource(url)
	if err != nil {
		return err
	}
	log.Println("Connected to the MySQL Server")
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DB = db
	return nil
}

func getDataSource(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to the mysql server")
	return db, nil
}

func Close() {
	DB.Close()
}