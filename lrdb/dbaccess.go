package lrdb

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func Test() {
	var name string
	db, err := sql.Open("sqlite3", "/Volumes/Claire/Photo/Photos 2015/Photos 2015-2.lrcat")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select name from AgLibraryCollection")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if (err != nil) {
			log.Fatal(err)
		}
		log.Println(name)
	}
}