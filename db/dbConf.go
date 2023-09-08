package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func DbAccess() {
	var err error
	Db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/zemus_api")
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		panic("err at ping: " + err.Error())
	}

}
