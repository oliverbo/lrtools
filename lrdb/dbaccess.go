package lrdb

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type RowLoader interface {
	Next() bool
	Scan(dest ...interface{}) error
}

func LoadData() {
	loadCollections()
}

func loadCollections() {

	db, err := sql.Open("sqlite3", "/Volumes/Claire/Photo/Photos 2015/Photos 2015-2.lrcat")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select id_local, name, parent from AgLibraryCollection")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	buildCollectionTree(rows)
}

func buildCollectionTree(rows RowLoader) {
	var name string
	var localId int64
	var parent sql.NullInt64

	for rows.Next() {
		err := rows.Scan(&localId, &name, &parent)
		if err != nil {
			log.Fatal(err)
		}

		// Initialize collection record
		c := GetCollectionById(localId)
		c.Name = name

		// Put it in the right place
		var p *Collection
		var parentId int64
		if parent.Valid {
			parentId = parent.Int64
		} else {
			parentId = CollectionRootId
		}

		p = GetCollectionById(parentId)
		p.AppendChild(c)
	}
}