package main

import (
	"github.com/oliverbo/lrtools/lrdb"
	"fmt"
)

func main() {
	lrdb.LoadData()

	root := lrdb.GetCollectionRoot()

	root.VisitChildren(func(c *lrdb.Collection) {
		fmt.Printf(c.Name)
	})
}
