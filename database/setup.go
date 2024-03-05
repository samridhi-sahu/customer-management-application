package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/samridhi-sahu/cms/protobuf"
)

var Client pb.CustomerServicesClient
var Db *sql.DB

func Setup() {
	var err error
	password := os.Getenv("MYSQL_PASSWORD")
	Db, err = sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatal("database connection error ", err)
	}
	err = Db.Ping()
	if err != nil {
		log.Println("connection viability error ", err)
	}
	log.Println("database connected")
}
