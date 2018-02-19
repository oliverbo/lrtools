package main

import (
	"github.com/oliverbo/lrtools/lrdb"
	"fmt"
	"flag"
	"os"
)

var cliParameters struct {
	dbPath string
	collectionPath string
	level int
	list bool
}

func main() {
	// Reading and writing current configuration
	lrdb.ReadConfig()
	defer lrdb.WriteConfig()

	parseCli()

	// Set catalog name
	if cliParameters.dbPath != "" {
		lrdb.Config.DbPath = cliParameters.dbPath
	}

	lrdb.LoadData()

	if cliParameters.list {
		listCollection()
	}
}

func parseCli() {
	flag.StringVar(&cliParameters.dbPath, "lrcat", "", "Path to Lightroom catalog")
	flag.StringVar(&cliParameters.collectionPath, "path", "", "Absolute path to collection")
	flag.IntVar(&cliParameters.level, "l", 1, "Number of levels to display")
	flag.BoolVar(&cliParameters.list, "list", false, "List collections in catalog")
	flag.Parse()
}

func listCollection() {
	var c *lrdb.Collection
	if cliParameters.collectionPath != "" {
		c = lrdb.FindCollectionByPath(cliParameters.collectionPath)
	} else {
		c = lrdb.GetCollectionRoot()
	}
	if c != nil {
		c.VisitChildren(cliParameters.level, false, func(l int, c *lrdb.Collection) {
			fmt.Println(c.Path())
		})
	} else {
		os.Stderr.WriteString("Invalid collecion root")
	}
}
