package main

import (
	"github.com/oliverbo/lrtools/lrdb"
	"fmt"
)

func main() {
	// Reading and writing current configuration
	lrdb.ReadConfig()
	defer lrdb.WriteConfig()

	lrdb.Config.DbPath = "/Volumes/Claire/Photo/Photos 2015/Photos 2015-2.lrcat"


	lrdb.LoadData()

	root := lrdb.GetCollectionRoot()

	root.VisitChildren(func(c *lrdb.Collection) {
		fmt.Println(c.Name)
	})
}
