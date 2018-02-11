package main

import (
	"github.com/oliverbo/lrtools/lrdb"
	"fmt"
)

func main() {
	lrdb.LoadData()

	root := lrdb.GetCollectionRoot()
	v := func(c *lrdb.Collection) {
		fmt.Printf(c.Name)
	}
	root.VisitChildren(v)
}
