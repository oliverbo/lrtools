package main

import (
	"github.com/oliverbo/lrtools/lrdb"
	"fmt"
	"flag"
)

var cliParameters struct {
	dbPath string
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
	flag.BoolVar(&cliParameters.list, "list", false, "List collections in catalog")
	flag.Parse()
}

func listCollection() {
	root := lrdb.GetCollectionRoot()

	root.VisitChildren(func(c *lrdb.Collection) {
		fmt.Println(c.Path())
	})
}
