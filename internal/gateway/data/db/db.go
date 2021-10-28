package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(url string) (*sql.DB, error) {
	if len(url) == 0 {
		user := viper.GetString("db.user")
		passwd := viper.GetString("db.passwd")
		host := viper.GetString("db.host")
		port := viper.GetInt("db.port")
		database := viper.GetString("db.database")
		url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, passwd, host, port, database)
	}
	db, err := getDataSource(url)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the MySQL Server")
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// DB = db
	return db, err
}

func getDataSource(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to the mysql server")
	return db, nil
}
