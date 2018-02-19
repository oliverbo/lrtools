package lrdb

import "testing"

func TestFindCollectionById(t *testing.T) {
	// Create new collection
	c := GetCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}


	// retrieve existing collection
	c.Name = "Test"
	c = GetCollectionById(5)
	if c == nil {
		t.Error("No collection pointer returned")
	}
	if c.localId != 5 || c.Name != "Test" {
		t.Errorf("Incorrect values %d, %s", c.localId, c.Name)
	}
}

func TestFindCollectionByPath(t *testing.T) {
	root = &Collection{localId:CollectionRootId, Name: "ROOT"}
	c1 := Collection{localId:1, Name: "Level 1"}
	c2 := Collection{localId:2, Name: "Level 2"}
	c3 := Collection{localId:3, Name: "Level 3"}
	c4 := Collection{localId:4, Name: "Level 4"}
	root.AppendChild(&c1)
	c1.AppendChild(&c2)
	c2.AppendChild(&c3)
	c3.AppendChild(&c4)

	c := GetCollectionByPath("/Level 1/Level 2")
	if c == nil {
		t.Error("No collection returned")
	} else if c != &c2 {
		t.Errorf("1: Incorrect collection found %d", c.localId)
	}
}

func TestCollection_FindChildByName(t *testing.T) {
	root = &Collection{localId:1, Name:"Parent"}
	c1 := Collection{localId:2, Name:"Child1"}
	c2 := Collection{localId:3, Name:"Child2"}
	c3 := Collection{localId:4, Name:"Child3"}
	root.AppendChild(&c1)
	root.AppendChild(&c2)
	root.AppendChild(&c3)

	cc := root.FindChildByName("Child2")
	if cc == nil {
		t.Error("1: No collection found")
	} else if cc != &c2 {
		t.Errorf("1: Incorrect collection found %d", cc.localId)
	}
}

func TestCollection_AppendChild(t *testing.T) {
	c := Collection {localId:5, Name:"Test"}
	child := Collection {localId:50, Name:"TestChild"}

	// Append first child
	c.AppendChild(&child)

	if c.first != &child || c.last != &child {
		t.Error("First Child: Parent collection not correctly wired")
	}
	if child.parent != &c || child.nextSibling != nil {
		t.Error("First Child: Child collection not correctly wired")
	}

	// Append second child
	child2 := Collection{localId:51, Name:"TestChild2"}
	c.AppendChild(&child2)
	if c.first != &child || c.last != &child2 {
		t.Error("2nd Child: Parent not correctly wired")
	}
	if child.nextSibling != &child2 {
		t.Error("2nd Child: First Child not correctly wired")
	}
	if child2.parent != &c || child2.nextSibling != nil {
		t.Error("2nd Child: 2nd child collection not correctly wired")
	}
}

func TestCollection_VisitChildren(t *testing.T) {
	c := Collection{localId:1, Name:"Parent"}
	c1 := Collection{localId:2, Name:"Child1"}
	c2 := Collection{localId:3, Name:"Child2"}
	c3 := Collection{localId:4, Name:"Child3"}
	c.AppendChild(&c1)
	c.AppendChild(&c2)
	c.AppendChild(&c3)

	var visited string
	f := func(c *Collection) {visited += c.Name}
	c.VisitChildren(f)

	if visited != "Child1Child2Child3" {
		t.Errorf("Vistor failed (%s)", visited)
	}
}

func TestCollection_Path(t *testing.T) {
	root = &Collection{localId:CollectionRootId, Name: "ROOT"}
	c1 := Collection{localId:1, Name: "Level 1"}
	c2 := Collection{localId:2, Name: "Level 2"}
	c3 := Collection{localId:3, Name: "Level 3"}
	c4 := Collection{localId:4, Name: "Level 4"}
	root.AppendChild(&c1)
	c1.AppendChild(&c2)
	c2.AppendChild(&c3)
	c3.AppendChild(&c4)

	path := c4.Path()

	if path != "/Level 1/Level 2/Level 3/Level 4" {
		t.Errorf("1: Incorrect collection path: '%s'", path)
	}

	path = c1.Path()
	if path != "/Level 1" {
		t.Errorf("2: Incorrect collection path: '%s'", path)
	}
}