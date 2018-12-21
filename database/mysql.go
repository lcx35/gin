package database

import (
	"database/sql"
	"log"
	"web/libs"

	_ "github.com/go-sql-driver/mysql"
)

var Conns *sql.DB

func init() {

	var err error
	username := libs.Conf.Read("mysql", "username")
	password := libs.Conf.Read("mysql", "password")
	dataname := libs.Conf.Read("mysql", "dataname")
	port := libs.Conf.Read("mysql", "port")
	host := libs.Conf.Read("mysql", "host")

	Driver := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"

	Conns, err = sql.Open("mysql", Driver)
	if err != nil {
		log.Fatal(err.Error())
	}

	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)

	err = Conns.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
