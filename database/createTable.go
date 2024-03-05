package database

import (
	"log"
)

func CreateTable() {
	query := `create table if not exists customer(
		id varchar(50) primary key,
		name varchar(50) not null,
		city varchar(50)
	)`
	_, err := Db.Query(query)
	if err != nil {
		log.Println("error in creating customer table", err)
	}
}
